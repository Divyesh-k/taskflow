# Task Manager

A REST API for task management built in Go. Clean layered architecture with JWT authentication, PostgreSQL persistence, and middleware-based request handling.

## What it does

Users can register, log in, and manage their tasks — create, read, update, delete. Each user only sees their own tasks. Authentication is handled via JWT tokens passed in request headers.

## Project structure
```
├── cmd/          # entry point
├── config/       # environment and database config
├── handlers/     # HTTP request handlers
├── middleware/   # JWT auth middleware
├── models/       # data structs
├── dto/          # request/response types
├── repository/   # database queries
├── services/     # business logic
├── routes/       # route definitions
└── utils/        # helpers
```
## Stack

| | |
|---|---|
| Language | Go |
| Router | standard or chi/gin (see go.mod) |
| Database | PostgreSQL |
| Auth | JWT |

## Running

```bash
# Set up environment
cp .env.example .env   # if present

# Run
go run cmd/main.go
```

## API

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /register | No | Create account |
| POST | /login | No | Get JWT token |
| GET | /tasks | Yes | List your tasks |
| POST | /tasks | Yes | Create a task |
| PUT | /tasks/:id | Yes | Update a task |
| DELETE | /tasks/:id | Yes | Delete a task |
