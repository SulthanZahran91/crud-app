# Backend

Go HTTP API backed by SQLite.

## Structure

```
main.go            — server entry point
handlers/todos.go  — HTTP request handling
services/todos.go  — business logic
repository/todos.go — database access
```

## Routes

| Method | Path       | Body                              | Response          |
|--------|------------|-----------------------------------|-------------------|
| GET    | /todos     | —                                 | `[]Todo`          |
| POST   | /todos     | `{"title":"..."}` | `Todo` (201)      |
| PUT    | /todos/:id | `{"title":"...","completed":bool}`| `Todo`            |
| DELETE | /todos/:id | —                                 | 204               |

## Build & Run

```bash
go build -o app .
./app
```

## Test

```bash
go test ./...
```
