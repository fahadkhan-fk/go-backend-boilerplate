# Go Backend Boilerplate 🚀

A clean, production-ready backend boilerplate built with **Golang**, featuring:

- ⚡️ [Fiber](https://github.com/gofiber/fiber) for fast HTTP routing
- 🛠 [GORM](https://gorm.io) ORM for PostgreSQL
- 🔐 JWT authentication with middleware
- 🧱 Clean architecture (handlers, services, repos)
- 🧾 Structured logging using [Zerolog](https://github.com/rs/zerolog)
- 🧪 Input validation with `validator.v10`
- 📦 Dependency injection & modular routing
- 📑 Full User CRUD support
- 🧵 Request ID middleware for traceability

---

## 📁 Folder Structure

```
.
├── config/              # Environment config loader
├── internal/
│   ├── db/              # GORM connection + migration
│   ├── models/          # GORM model definitions
│   ├── repo/            # Repositories (data access layer)
│   ├── service/         # Business logic
│   ├── handler/         # Fiber route handlers
│   ├── middleware/      # JWT, logging, request ID, etc.
│   ├── routes/          # Route setup + dependency injection
│   └── logger/          # Zerolog initialization
├── .env                 # Environment variables
├── go.mod
└── cmd/
    └── main.go         # Entry point
```

---

## 🧪 Features

- **/api/v1/register** – User registration with hashed password
- **/api/v1/login** – Login with JWT token response
- **/api/v1/users/me** – Authenticated user info
- **/api/v1/users** – Full CRUD endpoints
- **Request logging** with latency, IP, and status
- **Request ID injection** for trace correlation

---

## 🛠 Environment Variables

Example `.env`:

```env
SERVER_PORT=8080
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
DATABASE_DB=go_boilerplate
DATABASE_SSL_MODE=disable
JWT_SECRET=your_jwt_secret
```

---

## 📬 API Examples (cURL)

### 🔐 Register a User
```bash
curl --location 'http://localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"testuser@example.com", "password":"password123"}'
```

### 🔑 Login and Get JWT Token
```bash
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"testuser@example.com", "password":"password123"}'
```

### 👤 Get Authenticated User (`/users/me`)
```bash
curl --location 'http://localhost:8080/api/v1/users/me' \
--header 'Authorization: Bearer <your_jwt_token>'
```

### 📋 Get All Users
```bash
curl --location 'http://localhost:8080/api/v1/users' \
--header 'Authorization: Bearer <your_jwt_token>'
```

### 📄 Get a User by ID
```bash
curl --location 'http://localhost:8080/api/v1/users/1' \
--header 'Authorization: Bearer <your_jwt_token>'
```

### 🛠️ Update a User by ID
```bash
curl --location --request PATCH 'http://localhost:8080/api/v1/users/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <your_jwt_token>' \
--data-raw '{"email":"updateduser@example.com"}'
```

### ❌ Delete a User by ID
```bash
curl --location --request DELETE 'http://localhost:8080/api/v1/users/1' \
--header 'Authorization: Bearer <your_jwt_token>'
```

---

## 📦 Getting Started

```bash
cd cmd
go mod tidy
cp .env.example .env
go run main.go
```

---

## 🤝 Contributing

Want to improve or extend the boilerplate? PRs are welcome!

---

## 📜 License

MIT License

---

## 🙋‍♂️ Author

Developed by [Fahad Khan](https://pk.linkedin.com/in/fahadkhan-dev-engineer) — Backend Engineer focused on clean, scalable Go architecture.

Todo's:
- return err itself in the error responses
- db connection pool settings