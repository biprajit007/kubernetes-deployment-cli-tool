# Kubernetes Deployment CLI Tool

Small wrapper around `kubectl set image` with dry-run support.

## Features
- Namespace selection
- Dry-run by default
- Good base for deployment automation

## Usage
```bash
go run . --name api --image repo/api:latest --namespace prod --dry-run=false
```

## Disclaimer
This tool can change running workloads when dry-run is disabled.
