# AGENTS.md

## Project Overview

arr-dashboard is a simple Go web dashboard for monitoring systemd *arr services (Sonarr, Radarr, Lidarr, etc.).

## Running the Project

### Development
```bash
docker-compose up
```
Then visit http://localhost:8080

### Production Build
```bash
go build -o arr-dashboard ./cmd/server
./arr-dashboard
```

## Code Structure

```
/cmd/server/main.go       - HTTP server and embedded UI
/internal/config/         - Config loading and defaults
/internal/checker/       - Systemctl status checking
/internal/handler/       - HTTP API handlers
```

## Testing

- Run the server and check http://localhost:8080
- Verify config is created at ~/.arr-dashboard/config.json

## Conventions

- Go 1.21+
- Standard library only (no external dependencies)
- Clean separation: config → checker → handler → UI
