# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
APP_ENV=development
DATABASE_HOST=postgres-user
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=user_db
DATABASE_HOST_PORT=5435
DATABASE_DOCKER_PORT=5432
TIMEZONE=UTC
LOCAL_TIMEZONE=Asia/Bangkok
ADMIN_PASSWORD=banana
JWT_SECRET=NOVEMBER_RAIN
JWT_TOKEN_DURATION=24
REDIS_CONTAINER_NAME=redis-user
REDIS_HOST_PORT=6379
REDIS_DOCKER_PORT=6379
COMPOSE_PROJECT_NAME=demo-user-service
```

- For first time.

```bash
cd scripts && chmod +x dev-start.sh && ./dev-start.sh
```

- Later

```bash
cd scripts && ./dev-start.sh
```

## Database CLI

```bash
pgcli postgres://postgres:postgres@127.0.0.1:5435/user_db
```
