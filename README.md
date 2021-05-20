## Dependencies

- minio
- mysql

## Environments

```
APPLICATION_NAME (default: file-service)
```

```
PROFILE (default: dev)
```

```
PORT (default: 5050)
```

```
IS_CONNECT_CONFIG_SERVER (default: true)
```

```
CONFIG_SERVER_URL (default: http://localhost:8888/file-service/)
```

```
IS_CONNECT_SERVICE_REGISTRY (default: true)
```

```
DB_HOST
```

```
DB_PORT
```

```
DB_USERNAME
```

```
DB_PASSWORD
```

```
DB_NAME
```

```
DB_TYPE
```

```
MINIO_URL
```

```
MINIO_ACCESS_KEY
```

```
MINIO_SECRET_KEY
```

```
MINIO_BUCKET_NAME
```

```
MINIO_SECURE
```

## Docker Container Run

```
docker run -p 5050:5050 -e APPLICATION_NAME=file-service -e PROFILE=dev \
 -e PORT=5050 -e IS_CONNECT_CONFIG_SERVER=false -e IS_CONNECT_SERVICE_REGISTRY=false \
 -e DB_HOST=mysql -e DB_PORT=3306 -e DB_USERNAME=root -e DB_PASSWORD=admin \
 -e DB_NAME=kurban-be -e DB_TYPE=mysql -e MINIO_URL=mysql:9000 \
 -e MINIO_URL=minio:9000 -e MINIO_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE \
 -e MINIO_SECRET_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY \
 -e MINIO_BUCKET_NAME=kurban-be -e MINIO_SECURE=false --name file-service file-service
```