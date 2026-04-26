# Frontend

React SPA built with Vite.

## Components

```
src/
  App.jsx               — root component, state management, API calls
  api.js                — fetch wrappers for backend API
  components/
    TodoForm.jsx        — add new todo input
    TodoList.jsx        — list with inline edit, complete toggle, delete
```

## Dev

```bash
npm install
npm run dev
```

Proxies `/todos` to `http://localhost:8080` (configure in `vite.config.js`).

## Build

```bash
npm run build
```

Output in `dist/`.
