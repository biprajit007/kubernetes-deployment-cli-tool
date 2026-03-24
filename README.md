# kubernetes-deployment-cli-tool

A Go CLI wrapper around kubectl set image for updating a deployment container image, with dry-run behavior enabled by default.

## Key features

- Namespace, deployment name, and image flags
- Prints the exact kubectl command that would run
- Optional real execution when --dry-run=false

## Project structure

- `main.go` — CLI flag parsing and kubectl invocation
- `go.mod` — Go module definition

## Requirements

- Go 1.22+
- kubectl configured with cluster access

## Setup

```bash
git clone https://github.com/biprajit007/kubernetes-deployment-cli-tool.git
cd kubernetes-deployment-cli-tool
go build
```

## Usage

### Preview a rollout

```bash
go run . -namespace production -name web -image ghcr.io/example/web:v2
```

### Apply it

```bash
go run . -namespace production -name web -image ghcr.io/example/web:v2 -dry-run=false
```

## Safety notes

- Image updates affect live Kubernetes workloads. Verify the namespace, deployment name, and image tag before disabling dry-run.

## Limitations / next improvements

- Assumes the container name matches the deployment name
- No rollout status monitoring or rollback support
