# gochat

A lightweight Go chat server with WebSocket messaging, SQLite persistence, and a static frontend.

This project is designed as a simple full-stack chat application that is easy to run locally, easy to understand, and easy to extend. It combines a Go HTTP server, a WebSocket-based real-time messaging channel, and a small SQLite database for persistence.

The app is intentionally compact, but it still covers the core pieces you would expect in a practical chat system:

- user registration and login
- token-based authentication
- live messaging with WebSockets
- message history storage
- a browser-based frontend
- automatic database setup on startup

The goal of the repository is to provide a clear learning-friendly base for a chat product. You can use it as a starting point for experimenting with real-time features, backend architecture, or frontend integration.

## What it does

- Serves the UI from `static/` at `/`
- Exposes registration and login routes
- Supports real-time chat over `/ws`
- Stores chat history in `chatapp.db`
- Creates the database and tables automatically on startup
- Uses bcrypt to hash passwords before saving them
- Returns JWTs after successful login
- Keeps the app usable with minimal setup

## How it works

When the server starts, it initializes the SQLite database and ensures the required tables exist. After that, it sets up the HTTP routes and begins listening on the configured port.

The frontend is served directly from the `static` directory. That means you do not need a separate frontend build step just to run the project locally.

Registration stores a new user with a hashed password. Login checks the stored password hash and returns a JWT token if the credentials are valid.

The WebSocket endpoint is used for the live chat flow. Once the client connects, messages can be pushed in real time without polling the server repeatedly.

Message history is stored in SQLite, which keeps the app lightweight while still preserving earlier conversations across restarts.

## Requirements

- Go 1.24 or newer
- A terminal that can run `go`
- A browser for the frontend UI
- Permission to read and write files in the project folder

## Recommended Setup

The project works best when you run it from the repository root. That keeps the database file, compiled binary, and source code in the same place.

The expected folder structure looks like this:

```text
GO-CHAT/
  main.go
  auth.go
  db.go
  history.go
  user.go
  userlist.go
  ws.go
  models.go
  static/
  chatapp.db
  go.mod
  go.sum
  README.md
```

## Run It

From the project root:

```bash
go run .
```

By default the server listens on `8080`.

If you want to change the port in PowerShell:

```powershell
$env:PORT = '3000'
go run .
```

If `PORT` is not set, the app falls back to `8080`.

## Build It

```bash
go build -o gochat.exe .
```

On Windows, run the executable with:

```powershell
.\gochat.exe
```

Building the app is useful when you want a repeatable binary instead of running from source each time.

## Open In Browser

After the server starts, open:

```text
http://localhost:8080
```

If you changed the port, replace `8080` with the value you chose.

## Authentication

There is no default username or password in the codebase.

You must create your own account first using the registration route. After that, log in with the same credentials to receive a JWT token.

The registration flow works like this:

1. Send a username and password to `POST /register`
2. The server hashes the password with bcrypt
3. The hashed value is stored in SQLite
4. The server returns a success response

The login flow works like this:

1. Send the same username and password to `POST /login`
2. The server looks up the stored password hash
3. The provided password is compared with the hash
4. If the credentials are valid, the server returns a JWT

The token is generated in `auth.go` and is currently signed with a shared secret in code. For a production deployment, that secret should be moved to an environment variable.

### Example JSON Body

```json
{
  "username": "your_name",
  "password": "your_password"
}
```

### Example Registration Request

```bash
curl -X POST http://localhost:8080/register ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"alice\",\"password\":\"secret123\"}"
```

### Example Login Request

```bash
curl -X POST http://localhost:8080/login ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"alice\",\"password\":\"secret123\"}"
```

## API Routes

- `POST /register` - create a new user
- `POST /login` - log in and receive a JWT token
- `GET /ws` - WebSocket chat endpoint
- `GET /history` - fetch message history
- `GET /users` - list users
- `POST /deleteHistory` - delete message history

## Database

The app uses SQLite through the `chatapp.db` file in the project root.

The database is created automatically if it does not exist. The schema currently includes at least these tables:

- `users`
- `messages`

The `users` table stores usernames and password hashes.
The `messages` table stores sender, receiver, message content, and timestamps.

SQLite is a good fit here because it keeps the project self-contained. You do not need to install or configure a separate database server just to get started.

## Startup Behavior

When `main.go` starts the server, it does the following:

- initializes the database
- registers all HTTP handlers
- serves the static frontend
- wraps the mux with CORS middleware
- listens on the configured port

This means one command is enough to launch the full app locally.

## WebSocket Flow

The WebSocket route is the core of the live chat experience.

A client connects to `/ws`, and the server can then exchange messages with that client without the overhead of repeated HTTP polling.

That setup is useful for:

- instant message delivery
- low-latency updates
- a smoother chat experience
- reducing unnecessary requests

If you are extending the app, the WebSocket layer is the best place to add features like typing indicators, presence updates, or read receipts.

## Frontend Notes

The frontend lives in `static/` and is served directly by Go.

That means the UI can stay simple and still be tightly coupled to the backend. If you later split the frontend into a separate app, the backend routes can still stay the same.

Typical frontend tasks might include:

- logging in a user
- storing the returned token
- connecting to the WebSocket endpoint
- rendering new chat messages
- loading older messages from history

## Project Files

- `main.go` - server entry point and route setup
- `db.go` - database initialization and schema setup
- `ws.go` - WebSocket handling
- `auth.go` - authentication handlers
- `history.go` - chat history handlers
- `user.go` and `userlist.go` - user-related handlers
- `static/` - frontend assets
- `models.go` - shared data structures

## Development Notes

- `chatapp.db` is created in the project root if it does not already exist.
- The app enables CORS for local development.
- Passwords are never stored in plain text.
- The JWT secret is currently hardcoded and should be externalized for production.
- If you add new dependencies, run `go mod tidy`.
- If you change routes, update both the backend and the frontend together.
- If you change the database schema, test startup against a fresh `chatapp.db` file.

## Troubleshooting

If the app does not start, check the following:

- Confirm Go is installed and available in your terminal.
- Make sure you are running the command from the project root.
- Ensure the chosen port is not already in use.
- Check whether `chatapp.db` is writable in the project folder.
- Look for compile errors after adding new code or packages.

If login fails, verify that:

- the user was registered first
- the username matches exactly
- the password matches exactly
- the database file contains the expected account record

If the frontend loads but chat does not work, check the browser console and server logs for WebSocket errors.

## Security Notes

This repository is fine for local development and learning, but it should be improved before any production use.

Areas to improve before deployment include:

- move the JWT secret into an environment variable
- use stronger config management for credentials and secrets
- add better validation and error handling
- add rate limiting for auth endpoints
- add logging that is safe for production use
- consider HTTPS in front of the service
- review how message access and deletion should be authorized

## Why This Design

The stack is deliberately simple.

Go gives the backend fast startup and a small binary.
SQLite keeps the data layer local and easy to ship.
WebSockets provide the real-time part of chat without needing a heavier push system.

That combination makes the project good for:

- demonstrations
- classroom work
- experiments with real-time systems
- small internal tools
- first iterations of a chat feature

## Future Ideas

If you want to grow the project, here are some natural next steps:

- add private conversations
- add typing indicators
- add online presence
- add read receipts
- add message search
- add attachments
- add group chats
- add profile images
- add a richer frontend
- add tests for auth and chat behavior

## Quick Summary

Run the server with `go run .`, open `http://localhost:8080`, register a user, and then log in to receive a token.

From there, the WebSocket endpoint and SQLite-backed history give you the core pieces of a working chat app.
