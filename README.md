# Gobook

Simple web api to manage your list of books

## Getting started

run `go run cmd/main.go` to run the project

## Development

In order to kill anything on port 8080 run

```bash
alias nuke88="lsof -i tcp:8080 | grep LISTEN | awk '{print \$2}' | xargs kill"

nuke88
```
