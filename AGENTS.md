# CRUD App — React + Go + SQLite

## Repo-Specific Header

### Project Summary
A simple CRUD application with:
- React frontend
- Go backend API
- SQLite database
- Single binary deployment

### Non-Negotiables
- Keep it simple — no microservices
- SQLite for persistence (no external database)
- Single Go binary for backend
- React SPA for frontend

### Architecture Boundaries
- Frontend: React components for UI, API calls to backend
- Backend: Go handlers for HTTP routes, services for business logic, repository for DB access
- Database: SQLite file in `data/app.db`

### Key Commands
- Build backend: `cd backend && go build -o app`
- Build frontend: `cd frontend && npm run build`
- Run backend: `cd backend && ./app`
- Run frontend: `cd frontend && npm run dev`
- Test: `cd backend && go test ./...`

### Docs To Keep In Sync
- `README.md` — setup and usage
- `backend/README.md` — API endpoints
- `frontend/README.md` — component structure

### Repo-Specific Debugging Notes
- Check `backend/logs/` for API errors
- Use `sqlite3 data/app.db` to inspect DB directly
- Frontend: check browser console + network tab

---

## Baseline Rules
*(from ~/master-agents/MASTER_AGENTS.md)*

### Read Before Changing
- Read the relevant code and local docs before changing structure or behavior.
- Follow existing patterns before introducing new abstractions.
- Treat the codebase as the source of truth when docs drift.

### Decision Making
- Ask before making major architectural decisions when multiple valid directions exist.
- Do not introduce large structural changes without clear justification.
- Prefer the simplest approach that fits the current codebase.

### Change Discipline
- Keep changes scoped and atomic.
- Do not leave the repo in a knowingly broken state.
- Avoid unrelated edits unless they are required for correctness.

### Dependency Discipline
- Prefer the standard library and existing dependencies when they are sufficient.
- Do not add new dependencies without a concrete reason.
- Do not introduce abstraction layers speculatively.

### Code Quality
- Do not add redundant comments.
- Do not add noisy defensive code that does not match the surrounding codebase.
- Do not bypass type or validation problems with hacks when the issue can be fixed properly.
- Keep style consistent with the rest of the repo.

### Debugging And Verification
- Prefer terminal, logs, tests, and reproducible scripts for debugging.
- Verify changes with the relevant build, test, lint, or run commands when practical.
- If verification could not be completed, state that explicitly.

### Documentation
- When behavior, commands, routes, workflows, or structure change, update the relevant docs in the same turn.
- Do not defer doc updates if the current change already made them stale.

### Safety
- Do not run destructive actions without approval.
- Do not leak secrets, tokens, or private data.
- Do not perform external side effects unless explicitly requested.

---

## Verification Checklist
- Build passes, or the reason it was not run is stated explicitly.
- Tests were run when practical, or the gap is stated explicitly.
- Lint and format checks were run when relevant, or the gap is stated explicitly.
- Docs were updated if behavior, commands, routes, workflows, or structure changed.
- No unrelated changes were included without reason.
