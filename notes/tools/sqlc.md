## sqlc

A tool that auto writes Go repository layer based on SQL files.
Almost equivalent to Prisma but requires user to write the SQL code which is actually fine.

### Configuration

```bash
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
        sql_package: "pgx/v5"
        emit_json_tags: true
        emit_interface: true
```

When you add `emit_json_tags:true` you get structs generated with json tags:

```go
type User struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
}
```

Why `emit_interface` is a lifesaver

    With the interface: Your UserService can accept a Querier. Now, in your tests, you can create a "Mock" that satisfies that interface and returns fake data without ever touching Postgres.

### How to name methods in Go

Lean into "Grep-ability"
Instead of trying to make names short, make them descriptive. Here is how a Senior Go dev would name queries in sqlc:

| Node/Generic Style | Go/sqlc Style   | Why?                                  |
| ------------------ | --------------- | ------------------------------------- |
| findAll            | ListUsers       | Clear what is being listed.           |
| findById           | GetUserByID     | Specifies the lookup key.             |
| update             | UpdateUserEmail | If it only updates one field, say so! |
| create             | CreateUser      | Standard and explicit.                |
