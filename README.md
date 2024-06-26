# URL Shortener

A URL shortener service built with Go using the GIN framework, with storage options including MongoDB and Redis, and a frontend developed with HTML, CSS, and JavaScript.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Frontend](#frontend)
- [License](#license)

## Features

- Shorten long URLs
- Redirect short URLs to the original URL
- Storage using either MongoDB or Redis
- Frontend using HTML, CSS, and JavaScript

## Installation

### Prerequisites

- Go 1.22+
- MongoDB or Redis

### Clone the Repository

```bash
git clone https://github.com/BrainAxe/url-shortener.git

cd url-shortener
```

### Install Dependencies
```bash
go mod tidy
```

## Configuration
Create a .env file in the root directory with the following content:

```bash
MONGO_STORE_SOURCE=mongodb://localhost:27017
REDIS_STORE_SOURCE=localhost:6378
HOST_PORT=9000
```

### Run the Application
```bash
go run main.go
```
The application will start on `http://localhost:9000`.

## Usage

### API Endpoints
  - POST /api/shorten - Shorten a long URL
  - GET /api/:shortUrl - Redirect to the original URL

### Frontend
The frontend, built with HTML, CSS, and JavaScript, is accessible at the root endpoint.
 - Visit `http://localhost:9000` to access the frontend interface.

## License
This project is licensed under the MIT License. See `LICENSE` for more information.
