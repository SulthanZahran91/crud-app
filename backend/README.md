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

Environment variables:
- `PORT` — listen port, default `8080`
- `DB_PATH` — SQLite file path, default `data/app.db`

If `frontend/dist/index.html` exists, the backend also serves the built React app from `/` while keeping the API under `/todos`.

## Test

```bash
go test ./...
```
