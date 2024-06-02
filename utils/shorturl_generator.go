package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

/*
Why base-58 instead of standard base-64 encoding (Satoshi Nakamoto)?
- Don't want 0OIl characters that look the same in some fonts and could be used to create visually identical looking account numbers.
- A string with non-alphanumeric characters is not as easily accepted as an account number.
- Double-clicking selects the whole number as one word if it's all alphanumeric.
*/

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, strRandom string) string {
	urlHashBytes := sha256Of(initialLink + strRandom)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
