# Shortly — A Minimal URL Shortener API in Go

**Shortly** is a clean and minimalistic URL shortener built entirely in Go with SQLite as a local storage engine. This project was created as a learning exercise to understand how to:

- Accept and process data via HTTP
- Build a REST API in Go
- Handle persistent data with a relational database

> 📘 The focus was not on production readiness, but on building something clear and functional to learn Go’s `net/http`, `database/sql`, and API patterns.

---

## ✨ Features

- ✅ POST `/shorten` to create a short URL
- ✅ Optional expiration support (`expires_in_minutes`)
- ✅ Expired links return a 410 Gone status
- ✅ Expired links are automatically cleaned from the database
- ✅ Duplicate detection (same original URL returns the same short ID)
- ✅ GET `/{shortID}` to redirect users to the original URL
- ✅ SQLite-backed storage (no external dependencies)
- ✅ Fully local and lightweight

---

## 📦 Project Structure

```
shortly/
├── cmd/             # Server entrypoint
├── internal/
│   ├── api/         # HTTP handlers
│   ├── storage/     # SQLite database logic
│   └── utils/       # Short ID generator
├── go.mod
└── README.md
```

---

## 🧪 API Usage

### 🔗 Shorten a URL

**Request**
```bash
curl -X POST http://localhost:8080/shorten `
     -H "Content-Type: application/json" `
     -Body '{ "original_url": "https://example.com", "expires_in_minutes": 60 }'
```

**Response**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 🚀 Redirect

Visiting `http://localhost:8080/abc123` will redirect you to `https://example.com`.  
If the link has expired, you will receive a `410 Gone` response with a clear message.

---

## 🧼 Automatic Expiration Cleanup

Expired URLs are automatically removed:
- On server startup

---



## 🛠️ Getting Started

### 🔧 Requirements

- Go 1.21+
- No external DB or services

### 🧰 Run the app

```bash
go run ./cmd/server
```

---

## 💡 What I Learned

This project helped me learn how to:
- Build a REST API from scratch in Go
- Handle JSON input/output with validation
- Store and retrieve data with SQLite and SQL queries
- Implement time-based expiration and automatic cleanup
- Write modular and extensible backend code

---

## 🛠 Future Enhancements (Ideas)

- Track visit counts per URL
- Rate limiting or API keys
- Frontend UI for generating and managing links
- User authentication and per-user links
- Hosting and deployment via Docker

---

## 📄 License

MIT — open source and free to use for learning or personal projects.
