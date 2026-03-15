# Backlog

## Future Features

### Phase 2 (Nice to Have)

- [ ] Auto-discovery of *arr services via port scan
- [ ] HTTP health checks in addition to systemctl
  - Note: Would require making checker interface generic
- [ ] Service response time monitoring
- [ ] Dark/light theme toggle (currently dark only)
- [ ] Configurable port via config file or env var
- [ ] Service grouping (e.g., "*arr", "download clients")

### Phase 3 (Future)

- [ ] Docker container status support
- [ ] Push notifications on service down (webhook/Discord)
- [ ] Historical uptime tracking (requires database)
- [ ] Service restart button (systemctl restart)
- [ ] Multi-machine monitoring (central dashboard)

### Architecture Considerations

- For HTTP checks: Create a `Checker` interface in `internal/checker/` to allow both systemctl and HTTP implementations
- For multi-machine: Consider adding a `remote` package for fetching from other hosts
