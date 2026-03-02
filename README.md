# gochat

A simple Go-based chat server with WebSocket support and an embedded SQLite datastore. The server serves a static frontend from the `static` folder and exposes REST endpoints for registration, login, message history, and user lists.

**Features:**

- **WebSocket chat:** real-time messaging via `/ws`.
- **User auth endpoints:** `/register` and `/login`.
- **Message history:** retrieve and delete via `/history` and `/deleteHistory`.
- **Embedded SQLite DB:** uses `chatapp.db` in the project root.
- **Static frontend:** served from the `static` directory at `/`.

**Prerequisites:**

- Go 1.24 or newer (module support enabled).
- Git (to clone the repo).

**Quick Start (local)**

1. Clone the repository:

```bash
git clone https://github.com/kartikahlawat/gochat.git
cd gochat
```

2. Configure the port (optional):

- Edit `.env` and set `PORT = "8080"`, or set an environment variable:

PowerShell:

```powershell
$env:PORT = '8080'
go run .
```

CMD:

```cmd
set PORT=8080
go run .
```

Or simply run with the default port (8080):

```bash
go run .
```

3. Build (optional):

```bash
go build -o gochat .
# On Windows: .\gochat.exe
```

4. Open the app in your browser at `http://localhost:8080`.

**API Endpoints (server-side)**

- `POST /register` — register a new user
- `POST /login` — login and receive auth (implementation-specific)
- `GET/POST /ws` — WebSocket endpoint for real-time chat
- `GET /history` — message history
- `GET /users` — list users
- `POST /deleteHistory` — delete message history

**Database**

- The server uses `chatapp.db` (SQLite) in the project root. The DB and required tables are created automatically on startup.

**Notes**

- Static assets are served from the `static` directory.
- The server prints the absolute DB path and a small debug insert on startup for initial testing.
- Dependencies are managed via Go modules (`go.mod`). Run `go mod tidy` if you add packages.

**Development & Contribution**

- Feel free to open issues or PRs. If you plan to publish this repository, add a `LICENSE` file.

**Files of interest**

- `main.go` — server entrypoint and route registration
- `db.go` — DB initialization and schema
- `ws.go` — WebSocket handling
- `auth.go` — authentication handlers
- `static/` — frontend files served by the server

---

If you want, I can also:

- open the project in VS Code,
- run `go build` and start the server locally,
- or add a permissive `LICENSE` file before you publish.
