## go-migrate

a cli tool for database migrations.
I'm using flake.nix for installing it to my projects.

```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migrations <migration_name>
```

`-seq`: sequential name of files -> 001 002 003
`-ext`: file extension -> ex: `sql`
`-dir`: directory to create the migrations at.
