# Python Web Development

## Flask Application
```bash
docker run -p 5000:5000 -v $(pwd):/app public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest python app.py
```

## Django Application
```bash
docker run -p 8000:8000 -v $(pwd):/app public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest python manage.py runserver 0.0.0.0:8000
```

## FastAPI Application
```bash
docker run -p 8000:8000 -v $(pwd):/app public.ecr.aws/registry/{{REPOSITORY_NAME}}:latest uvicorn main:app --host 0.0.0.0 --port 8000
```
