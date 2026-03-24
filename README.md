# mini-load-balancer

A very small Go reverse proxy that round-robins incoming requests between two hard-coded backend URLs.

## Key features

- Round-robin balancing using an atomic counter
- Built on Go's standard reverse proxy

## Project structure

- `main.go` — Proxy and backend selection logic
- `go.mod` — Go module definition

## Requirements

- Go 1.22+
- Two backend services listening on localhost:9001 and localhost:9002 unless you edit the code

## Setup

```bash
git clone https://github.com/biprajit007/mini-load-balancer.git
cd mini-load-balancer
go run .
```

## Usage

### Send a request through the proxy

```bash
curl http://localhost:8083/
```

## Safety notes

- Backends are hard-coded. Review and change them before using this outside a local test environment.

## Limitations / next improvements

- No health checks, sticky sessions, retries, or configuration file
- Only two local HTTP backends are defined in code
