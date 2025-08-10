# Usage Guide

## Basic Spring Boot Application

```bash
# Run a basic Spring Boot application
docker run -p 8080:8080 \
  -v $(pwd)/app.jar:/app/application.jar \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Production Deployment

```bash
# Production deployment with environment configuration
docker run -d \
  --name java-app \
  --restart unless-stopped \
  -p 8080:8080 \
  -e SPRING_PROFILES_ACTIVE=production \
  -e JAVA_OPTS="-Xmx2g -Xms1g" \
  -e DATABASE_URL="jdbc:postgresql://db:5432/myapp" \
  -v /app/logs:/opt/app/logs \
  -v $(pwd)/application.jar:/app/application.jar \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Docker Compose Setup

```yaml
version: '3.8'
services:
  app:
    image: public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
    ports:
      - "8080:8080"
    environment:
      - SPRING_PROFILES_ACTIVE=production
      - DATABASE_URL=jdbc:postgresql://db:5432/myapp
      - REDIS_URL=redis://redis:6379
    volumes:
      - ./app.jar:/app/application.jar
      - ./logs:/opt/app/logs
    depends_on:
      - db
      - redis

  db:
    image: postgres:15
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password

  redis:
    image: redis:7-alpine
```

## Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: java-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: java-app
  template:
    metadata:
      labels:
        app: java-app
    spec:
      containers:
      - name: app
        image: public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
        ports:
        - containerPort: 8080
        env:
        - name: SPRING_PROFILES_ACTIVE
          value: "kubernetes"
        - name: JAVA_OPTS
          value: "-Xmx2g -Xms1g -XX:+UseG1GC"
        resources:
          requests:
            memory: "1Gi"
            cpu: "500m"
          limits:
            memory: "3Gi"
            cpu: "2"
        livenessProbe:
          httpGet:
            path: /actuator/health
            port: 8080
          initialDelaySeconds: 60
          periodSeconds: 30
        readinessProbe:
          httpGet:
            path: /actuator/health/readiness
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        volumeMounts:
        - name: app-jar
          mountPath: /app/application.jar
          subPath: application.jar
      volumes:
      - name: app-jar
        configMap:
          name: app-config
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `SPRING_PROFILES_ACTIVE` | Active Spring profiles | default | No |
| `JAVA_OPTS` | JVM options | -Xmx1g -Xms512m | No |
| `DATABASE_URL` | Database connection URL | - | Yes |
| `REDIS_URL` | Redis connection URL | - | No |
| `LOG_LEVEL` | Application log level | INFO | No |

### Health Monitoring

```bash
# Application health
curl http://localhost:8080/actuator/health

# Readiness check
curl http://localhost:8080/actuator/health/readiness

# Metrics endpoint
curl http://localhost:8080/actuator/metrics
```

## Security Best Practices

### Non-root User
```bash
# Verify non-root execution
docker exec container-name id
# Output: uid=1001(appuser) gid=1001(appuser)
```

### Vulnerability Scanning
```bash
# Scan for vulnerabilities
docker scout cves public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

## Troubleshooting

### Memory Issues
```bash
# Increase heap size
docker run -e JAVA_OPTS="-Xmx4g -Xms2g" public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

### Debug Mode
```bash
# Enable debug logging
docker run -e LOG_LEVEL=DEBUG public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```

### JVM Analysis
```bash
# Enable JVM debugging
docker run -e JAVA_OPTS="-XX:+PrintGCDetails -XX:+PrintGCTimeStamps" \
  public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest
```
