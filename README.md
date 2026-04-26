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
- `DB_PATH` — SQLite file path (default: `data/app.db`)
- `PORT` — listen port (default: `8080`)

When `frontend/dist` exists, the Go server also serves the built React app on `/`.

### Frontend

```bash
cd frontend
npm install
npm run build
npm run dev
```

Runs on `http://localhost:5173`. During local development, Vite proxies `/todos` to backend.

## Docker / Dokploy

This repo includes a multi-stage `Dockerfile` for Dokploy deployment.

- builds the React frontend
- builds the Go backend
- serves the frontend and API from one container
- persists SQLite at `/app/data/app.db`

Container defaults:
- `PORT=8080`
- `DB_PATH=/app/data/app.db`

In Dokploy, deploy it as a Dockerfile app and expose port `8080`.

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
