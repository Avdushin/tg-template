# Шаблон Telegram бота

## Запуск

1. Локально:

```bash
docker compose up -d db
make run
```

2. Через Docker:

```bash
make run MODE=docker
```

3. Через docker-compose:

```bash
make run MODE=compose
```

4. С помощью [Air]():
```bash
docker compose up -d db
make air
```