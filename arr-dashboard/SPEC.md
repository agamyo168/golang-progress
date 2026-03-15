# arr-dashboard Specification

## Project Overview

**Project Name**: arr-dashboard  
**Type**: Simple web dashboard for monitoring *arr services  
**Core Functionality**: Display up/down status of systemd *arr services with clickable links  
**Target Users**: Home lab users running Sonarr, Radarr, Lidarr, etc.

## Goals

- [x] Single Go binary with embedded static frontend
- [x] Config file at `~/.arr-dashboard/config.json`
- [x] Auto-generate default config on first run
- [x] Display service name, status (up/down), and link to web UI
- [x] Manual refresh button to re-check status
- [x] Clean, minimal dark theme UI

## Non-Goals

- Auto-discovery of services
- Historical status tracking
- Push notifications
- Docker container monitoring

## Architecture

```
┌─────────────────────────────────────┐
│           Go HTTP Server            │
│  ┌─────────────┐  ┌──────────────┐  │
│  │   /api/*    │  │   Static UI  │  │
│  │  Handlers   │  │   (HTML/JS)  │  │
│  └──────┬──────┘  └──────────────┘  │
│         │                           │
│  ┌──────▼──────┐                    │
│  │  Service    │                    │
│  │  Checker    │                    │
│  └──────┬──────┘                    │
│         │                           │
│  ┌──────▼──────┐                    │
│  │ systemctl   │                    │
│  │   exec      │                    │
│  └─────────────┘                    │
└─────────────────────────────────────┘
```

## Configuration

**Location**: `~/.arr-dashboard/config.json`

**Schema**:
```json
{
  "host": "192.168.1.100",
  "services": [
    { "systemd": "sonarr",  "display": "Sonarr",  "port": 8989 },
    { "systemd": "radarr",  "display": "Radarr",  "port": 7878 },
    { "systemd": "lidarr",  "display": "Lidarr",  "port": 8686 },
    { "systemd": "readarr", "display": "Readarr", "port": 8787 },
    { "systemd": "prowlarr","display": "Prowlarr","port": 9696 }
  ]
}
```

**Defaults**:
- `host`: auto-detect local IP, fallback to `localhost`
- `services`: all 5 common *arr services on standard ports

## API Design

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/services` | GET | Returns all services with status |
| `/api/refresh` | POST | Triggers fresh systemctl check |
| `/` | GET | Dashboard HTML |

**Response Format** (`/api/services`):
```json
{
  "services": [
    {
      "name": "Sonarr",
      "systemd": "sonarr",
      "port": 8989,
      "url": "http://192.168.1.100:8989",
      "status": "active"
    }
  ],
  "lastRefresh": "2026-03-14T12:00:00Z"
}
```

## UI Specification

- Single page dashboard
- Header with title "arr-dashboard" and refresh button
- Grid or list of service cards
- Each card shows:
  - Service name (display field)
  - Status indicator (green dot = active, red dot = inactive)
  - Clickable link to service web UI
- Dark theme

## Deployment

- Default port: 8080
- Serve on all interfaces (`0.0.0.0:8080`)
