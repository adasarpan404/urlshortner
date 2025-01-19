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
