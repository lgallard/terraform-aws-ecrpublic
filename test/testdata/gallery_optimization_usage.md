# Usage Instructions

## Quick Start

```bash
# Pull the latest image
docker pull public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest

# Run with default configuration
docker run -p 3000:3000 public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Production Usage

```bash
# Production deployment
docker run -d \
  --name {{REPOSITORY_NAME}} \
  --restart unless-stopped \
  -p 3000:3000 \
  -e NODE_ENV=production \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{REPOSITORY_NAME}}
spec:
  replicas: 3
  selector:
    matchLabels:
      app: {{REPOSITORY_NAME}}
  template:
    metadata:
      labels:
        app: {{REPOSITORY_NAME}}
    spec:
      containers:
      - name: app
        image: public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
        ports:
        - containerPort: 3000
        livenessProbe:
          httpGet:
            path: /health
            port: 3000
```

## Environment Variables

- `NODE_ENV`: Environment mode (development/production)
- `PORT`: Application port (default: 3000)
- `LOG_LEVEL`: Logging level (debug/info/warn/error)

## Health Checks

```bash
curl http://localhost:3000/health
```
