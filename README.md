# gochat

A lightweight Go chat server with WebSocket messaging, SQLite persistence, and a static frontend.

## What it does

- Serves the UI from `static/` at `/`
- Exposes auth routes for registration and login
- Supports real-time chat over `/ws`
- Stores chat history in `chatapp.db`
- Creates the database and tables automatically on startup

## Requirements

- Go 1.24 or newer
- Windows PowerShell, Terminal, or any shell that can run `go`

## Run It

From the project root:

```bash
go run .
```

By default the server listens on `8080`.

To change the port in PowerShell:

```powershell
$env:PORT = '3000'
go run .
```

## Build It

```bash
go build -o gochat .
```

On Windows, run the executable with:

```powershell
.\gochat.exe
```

## Open In Browser

After the server starts, open:

```text
http://localhost:8080
```

## API Routes

- `POST /register`
- `POST /login`
- `GET /ws`
- `GET /history`
- `GET /users`
- `POST /deleteHistory`

## Project Files

- `main.go` - server entry point and route setup
- `db.go` - database initialization and schema setup
- `ws.go` - WebSocket handling
- `auth.go` - authentication handlers
- `history.go` - chat history handlers
- `user.go` and `userlist.go` - user-related handlers
- `static/` - frontend assets

## Notes

- `chatapp.db` is created in the project root if it does not already exist.
- The app enables CORS for local development.
- If you add new dependencies, run `go mod tidy`.
