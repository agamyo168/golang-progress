# Backlog

## Completed

- [x] Service selection UI (toggle on/off from browser)
- [x] Catppuccin styling
- [x] Config file permissions fix (0600)
- [x] Auto-detect local IP
- [x] Fixed JavaScript template literal syntax

## Future Features

### Phase 2 (Nice to Have)

- [ ] Auto-discovery of *arr services via port scan
- [ ] HTTP health checks in addition to systemctl
  - Note: Would require making checker interface generic
- [ ] Service response time monitoring
- [ ] Dark/light theme toggle (currently dark only)
- [ ] Configurable port via config file or env var
- [ ] Service grouping (e.g., "*arr", "download clients")
- [ ] systemd-nspawn container monitoring (for services running in containers)

### Phase 3 (Future)

- [ ] Docker container status support
- [ ] Push notifications on service down (webhook/Discord)
- [ ] Historical uptime tracking (requires database)
- [ ] Service restart button (systemctl restart)
- [ ] Multi-machine monitoring (central dashboard)

### Nix Packaging (Future)

- Consider adding Nix package to flake.nix using `buildGoModule`
- Could create separate nix-config repo to import this as input
- Use Renovate or similar for automated version updates

## Architecture Considerations

- For HTTP checks: Create a `Checker` interface in `internal/checker/` to allow both systemctl and HTTP implementations
- For multi-machine: Consider adding a `remote` package for fetching from other hosts
