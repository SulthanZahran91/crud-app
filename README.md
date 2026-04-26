# CRUD App

React + Go + SQLite todo application.

## Setup

### Backend

```bash
cd backend
go build -o app
./app
```

Runs on `http://localhost:8080`.

Environment variables:
- `DB_PATH` — SQLite file path (default: `../data/app.db`)
- `PORT` — listen port (default: `8080`)

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Runs on `http://localhost:5173`. Proxies `/todos` to backend.

## API

| Method | Path          | Description     |
|--------|---------------|-----------------|
| GET    | /todos        | List all todos  |
| POST   | /todos        | Create todo     |
| PUT    | /todos/:id    | Update todo     |
| DELETE | /todos/:id    | Delete todo     |

## Tests

```bash
cd backend && go test ./...
```
