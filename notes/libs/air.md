## Air

it's like running nodemon in node.js but instead it's for Golang.
Allows hot-reload on save.

### Install

```bash
go install github.com/air-verse/air@latest
# Make sure that you have path for golang binaries if not the default is ~/go/bin
# Add ~/go/bin to your $PATH
```

### Usage

```bash
# Initializes a TOML file (another config file.)
air init .
# after you configure the TOML file
air # Will automatically watch for changes.

```
