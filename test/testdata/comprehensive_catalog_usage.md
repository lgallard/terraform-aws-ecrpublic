# Usage Guide

## Quick Start

### Basic Deployment
```bash
# Pull the latest stable release
docker pull public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest

# Run with default configuration
docker run -d \
  --name my-app \
  -p 8080:8080 \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

### Production Deployment
```bash
# Production deployment with custom configuration
docker run -d \
  --name production-app \
  --restart unless-stopped \
  -p 8080:8080 \
  -e NODE_ENV=production \
  -e LOG_LEVEL=info \
  -e METRICS_ENABLED=true \
  -v /var/log/app:/app/logs \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Configuration Options

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `NODE_ENV` | Environment mode | development | No |
| `PORT` | Application port | 8080 | No |
| `LOG_LEVEL` | Logging level | info | No |
| `METRICS_ENABLED` | Enable metrics | false | No |
| `DATABASE_URL` | Database connection | - | Yes |
| `REDIS_URL` | Redis connection | - | No |

### Volume Mounts

- `/app/logs`: Application log files
- `/app/config`: Configuration files
- `/app/data`: Persistent data storage

## Health Monitoring

### Health Check Endpoints

```bash
# Application health
curl http://localhost:8080/health

# Readiness check
curl http://localhost:8080/ready

# Metrics endpoint
curl http://localhost:8080/metrics
```

### Expected Responses

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "uptime": 3600,
  "version": "1.0.0"
}
```

## Kubernetes Deployment

### Basic Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: enterprise-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: enterprise-app
  template:
    metadata:
      labels:
        app: enterprise-app
    spec:
      containers:
      - name: app
        image: public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
        ports:
        - containerPort: 8080
        env:
        - name: NODE_ENV
          value: "production"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

## Security Considerations

### Running as Non-root
This container runs as a non-root user by default:

```bash
# Verify non-root execution
docker run --rm public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest id
# Output: uid=1001(appuser) gid=1001(appuser) groups=1001(appuser)
```

### Scanning for Vulnerabilities
```bash
# Scan image for vulnerabilities
docker scout cves public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Troubleshooting

### Common Issues

**Port Already in Use**
```bash
# Use different port
docker run -p 8081:8080 public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

**Permission Issues**
```bash
# Check container user
docker run --rm public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest whoami
```

### Debug Mode
```bash
# Run in debug mode
docker run -e LOG_LEVEL=debug public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Support & Contributing

- **Issues**: Report issues on the project repository
- **Documentation**: Comprehensive docs at project homepage
- **Security**: Report security issues privately
- **Contributing**: Pull requests welcome with proper testing
