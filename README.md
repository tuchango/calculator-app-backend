# Calculator App Backend

Backend для приложения калькулятора на Go с использованием Echo, GORM и PostgreSQL.

## Описание

Этот проект реализует REST API для хранения и вычисления математических выражений. Все вычисления сохраняются в базе данных PostgreSQL.

## Стек технологий

- `Go` 1.24.6
- `Echo` (HTTP сервер)
- `GORM` (ORM)
- `PostgreSQL`
- `govaluate` (вычисление выражений)
- `godotenv` (загрузка переменных окружения)
- `uuid` (генерация уникальных идентификаторов)

## Запуск

1. Скопируйте `.env-example` в `.env` и укажите параметры подключения к базе данных.
2. Убедитесь, что PostgreSQL запущен и доступен.
3. Установите зависимости:

    ```sh
    go mod tidy
    ```

4. Запустите сервер:

    ```sh
    go run main.go
    ```

Сервер будет доступен на `localhost:8080`.

## API

Примеры запросов находятся в [api.http](api.http).

### Получить все вычисления

```
GET /calculations
```

### Создать новое вычисление

```
POST /calculations
{
    "expression": "12+34"
}
```

### Изменить вычисление

```
PATCH /calculations/:id
{
    "expression": "1111+3"
}
```

### Удалить вычисление

```
DELETE /calculations/:id
```

## Структура проекта

- `main.go` — точка входа
- `internal/calculation/` — бизнес-логика (handler, service, repository, model)
- `internal/db/` — инициализация базы

## Тестирование

Проект включает базовые API тесты. Подробнее о тестах можно узнать в [tests/README.md](tests/README.md).

Для запуска тестов выполните:

```sh
go test ./tests
```