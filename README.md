# URL Shortener (Go)

A minimal URL shortener service built using **Go and net/http**.
It supports creating short URLs, fetching URL details, and redirecting users while tracking visit counts.
Data is stored in memory for learning purposes.

---

## Features

1. Create a short URL from an original URL
2. Redirect using `/r/{id}`
3. Track visit count
4. Health check endpoint
5. JSON based API

---

## Tech Stack

1. Go
2. net/http
3. In memory storage (map)

---
```
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── internal
│   ├── handler
│   │   ├── health.go
│   │   ├── redirect.go
│   │   └── shorten.go
│   ├── model
│   │   └── url.go
│   ├── repository
│   │   └── memory.go
│   └── service
│       └── shortener.go
├── main.go
└── README.md
```


## API Endpoints

### Health Check

```
GET /
```

Response:

```json
{
  "status": "ok"
}
```

---

### Create Short URL

```
POST /shorten
```

Request body:

```json
{
  "original_url": "https://github.com/aialok"
}
```

Response:

```json
{
  "id": "a1b2c3",
  "original_url": "https://github.com/aialok",
  "short_url": "a1b2c3",
  "visits": 0,
  "created_at": "2025-01-01T10:00:00Z"
}
```

---

### Get URL Details

```
GET /url?short_url=a1b2c3
```

Response:

```json
{
  "id": "a1b2c3",
  "original_url": "https://github.com/aialok",
  "short_url": "a1b2c3",
  "visits": 1,
  "created_at": "2025-01-01T10:00:00Z"
}
```

---

### Redirect to Original URL

```
GET /r/a1b2c3
```

Redirects to the original URL and increments visit count.

---

## How to Run

1. Clone the repository
2. Run the server

```bash
go run .
```

Server starts at:

```
http://localhost:3000
```

---

## Notes

1. This project uses in memory storage and is not persistent
2. Hash based short URLs may collide in rare cases
3. Built for learning Go, HTTP servers, and basic system design

---

## Next Improvements

1. Persistent database
2. Base62 ID generation
3. Concurrency safety with mutex
4. Tests
5. Router based path parameters


