package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UrlDoc struct {
	ID        primitive.ObjectID `bson:"_id"`
	LongUrl   string             `bson:"longUrl"`
	ShortUrl  string             `bson:"shortUrl"`
	CreatedAt time.Time          `bson:"createdAt"`
}

// Top level declarations
var (
	ctx          = context.Background()
	StoreService = &StorageService{}
)

const CacheDuration = 6 * time.Hour

// Database interface
type StorageStrategy interface {
	SaveUrlMapping(string, string)
	RetrieveInitialUrl(string) string
}

// Define the struct wrapper around raw Mongo client
type MongoStorageService struct {
	mongoClient *mongo.Client
}

func (ss *MongoStorageService) SaveUrlMapping(shortUrl string, originalUrl string) {
	collection := ss.mongoClient.Database("test").Collection("urls")
	newDoc := &UrlDoc{
		ID:        primitive.NewObjectID(),
		LongUrl:   originalUrl,
		ShortUrl:  shortUrl,
		CreatedAt: time.Now(),
	}
	_, err := collection.InsertOne(context.TODO(), newDoc)

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func (ss *MongoStorageService) RetrieveInitialUrl(shortUrl string) string {
	var result bson.M
	collection := ss.mongoClient.Database("test").Collection("urls")
	err := collection.FindOne(context.Background(), bson.D{{Key: "shortUrl", Value: shortUrl}}).Decode(&result)

	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	longUrl := fmt.Sprint(result["longUrl"]) // converting to string
	return longUrl
}

// Define the struct wrapper around raw Redis client
type RedisStorageService struct {
	redisClient *redis.Client
}

func (ss *RedisStorageService) SaveUrlMapping(shortUrl string, originalUrl string) {
	err := ss.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func (ss *RedisStorageService) RetrieveInitialUrl(shortUrl string) string {
	result, err := ss.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}

// StorageService struct now holds the storage strategy
type StorageService struct {
	Strategy StorageStrategy
}

func InitializeStore(storageType string) *StorageService {
	var strategy StorageStrategy
	switch storageType {
	case "mongo":
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_STORE_SOURCE")))
		if err != nil {
			panic(fmt.Sprintf("Error init Mongo: %v", err))
		}
		fmt.Printf("\n MongoDB connected\n")
		strategy = &MongoStorageService{client}

	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_STORE_SOURCE"),
			Password: "",
			DB:       0,
		})

		pong, err := client.Ping(ctx).Result()
		if err != nil {
			panic(fmt.Sprintf("Error init Redis: %v", err))
		}

		fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
		strategy = &RedisStorageService{client}
	default:
		panic("Invalid storage type")
	}
	StoreService.Strategy = strategy
	return StoreService
}
