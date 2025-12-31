## Tools

### Starting

You’re used to package.json and node_modules. In Go, we use Go Modules. package.json ≈ go.mod: This file tracks your dependencies and Go version.
node_modules ≈ Module Cache: Go doesn't put a giant folder in every project. It keeps a global cache on your machine, which makes builds much faster.

```bash
npm install # adds packages to node_modules
go get # adds packages to go.mod
```

### Init

```bash
# init a project with go.mod
go mod init github.com/agamyo168/social-blog
```

### Go run

```bash
# runs main.go without compiling it.
go run main.go

# if multiple files
go run *.go
```
