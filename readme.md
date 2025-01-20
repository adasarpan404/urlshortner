# URL SHORTNER

A simple URL Shortener built using Go, GORM, and SQLite. This application allows users to generate short URLs for any given long URL and provides a redirect feature to the original URL when the short URL is accessed.

## Features

- Generate short URLs for long URLs.
- Persist data using SQLite with GORM.
- Redirect short URLs to the original URL.

## Requirements

- Go 1.20+
- SQLite
- (Optional) GCC (required if using go-sqlite3 with CGO).

## Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### Install Dependencies

```bash
go mod tidy
```

### Run the application

```bash
go run main.go
```

## Endpoints

### 1. Shorten URL

**POST** `/shorten`

Request Body:

```json
{
  "url": "https://www.example.com"
}
```

### 2. Redirect to Original URL

**GET** /{shortcode}

- Replace {shortCode} with the short code generated from the /shorten endpoint.

- The service will redirect to the original URL.

## How It Works

### 1. Generate a short url

- A random 6-character alphanumeric string is generated for the short URL.

- The mapping between the short URL and the original URL is saved in the SQLite database.

### 2. Redirect To The Original URL

- When a user accesses the short URL, the application queries the database for the original URL and redirects the user.

## Configuration

If using go-sqlite3 and encountering CGO issues, you can:

1. Enable CGO

   ```bash
   export CGO_ENABLED=1
   ```

2. Switch to a pure Go SQLite driver:

   ```bash
   go get modernc.org/sqlite
   ```

## Future Enhancements

- Create authentication

- Add expiration for short URLs.

- Allow users to create custom short codes.

- Implement analytics for tracking clicks.

- Add user authentication for managing URLs.
