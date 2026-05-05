# go-todo

[![Go](https://img.shields.io/badge/Go-1.26.2-blue)](https://golang.google.cn)
[![gin](https://img.shields.io/badge/gin-1.12.0-brightgreen)]()
[![MySQL](https://img.shields.io/badge/MySQL-8.0.23-blue)](https://www.mysql.com)

A demo project to practice Golang.



## Project Structure

```
go-todo/
├── api/            # API layer (Gin HTTP server)
├── service/        # Business logic layer
├── repository/     # Data access layer
├── common/         # Common utilities
├── db/             # Database schema
└── requests/       # API request examples
```



## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: MySQL



## Quick Start

### 1. Database Setup

```bash
mysql -u root -p < db/init.sql
```


### 2. Configure Database Connection

Edit `repository/todo_repository.go` to update database credentials:

```go
const (
	dbUser     = "root"
	dbPassword = "your_password"
	dbName     = "my_demo"
	dbAddr     = "localhost:3306"
)
```


### 3. Run

```bash
cd api && go run api.go
```

Server runs at `http://127.0.0.1:8000`



## API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/ping` | Health check |
| POST | `/todos` | Create todo |
| PUT | `/todos/:todoId` | Update todo |
| GET | `/todos/:todoId` | Get todo by ID |
| DELETE | `/todos/:todoId` | Delete todo |
| GET | `/todos?category=xxx` | Query by category |
| GET | `/todos?content=xxx` | Query by content |
| GET | `/todos?isComplete=Y\|N` | Query by completion status |
| GET | `/todos/categories` | Query category |


### Examples

**Create todo:**
```bash
curl -X POST http://localhost:8000/todos \
  -H "Content-Type: application/json" \
  -d '{"content":"学习Go","category":"工作","isComplete":"N"}'
```

**Query by isComplete:**
```bash
curl "http://localhost:8000/todos?isComplete=N"
```

**Delete todo:**
```bash
curl -X DELETE http://localhost:8000/todos/1
```

**Update todo:**
```bash
curl -X PUT --json '{
        "content": "取快递📦",
        "category": "日常生活",
        "isComplete": "Y",
        "deadline": "2026-05-05T13:35:00+08:00"
}' http://localhost:8000/todos/2
```

**Query categories:**
```bash
curl http://localhost:8000/todos/categories
```



## Data Model

| Field | Type | Description |
|-------|------|-------------|
| todo_id | int | Primary key |
| content | varchar(255) | Todo content |
| category | varchar(50) | Category |
| is_complete | char(1) | Y - completed, N - not completed |
| deadline | datetime | Deadline |
| create_at | timestamp | Creation time |
| update_at | timestamp | Last update time |
