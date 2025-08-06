# People API

**REST API для управления людьми с автоматическим получением данных из внешних сервисов**

[![Go Version](https://img.shields.io/badge/Go-1.24.3-blue.svg)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-green.svg)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Описание

People API - это современное веб-приложение на Go, которое предоставляет REST API для управления информацией о людях. Система автоматически получает недостающие данные (возраст, пол, национальность) из внешних API сервисов.

### Ключевые возможности

- 🔍 **Автоматическое получение данных** из внешних API (agify.io, genderize.io, nationalize.io)
- 👥 **Управление людьми** - создание, поиск, обновление
- 📧 **Управление email адресами** для каждого человека
- 👫 **Система друзей** - добавление и просмотр дружеских связей
- 🗄️ **PostgreSQL** с автоматическими миграциями
- 📚 **Swagger UI** для интерактивной документации API
- 🐳 **Docker** поддержка для простого развертывания
- 🔧 **Чистая архитектура** с разделением на слои

## Архитектура

Проект использует чистую архитектуру с четким разделением на слои:

```
┌─────────────────┐
│   HTTP Layer    │  ← Router, Server
├─────────────────┤
│  Handler Layer  │  ← API Handlers
├─────────────────┤
│  Service Layer  │  ← Business Logic
├─────────────────┤
│ Repository Layer│  ← Data Access
├─────────────────┤
│  Database Layer │  ← PostgreSQL
└─────────────────┘
```

### Структура проекта

```
├── cmd/                   # Точка входа приложения
│   └── app/               # Основное приложение
│       ├── main.go        # main
│       └── di.go          # Dependency Injection
├── internal/              # Внутренние пакеты
│   ├── api/               # API обработчики
│   │   └── people/v1/     # People API v1
│   ├── external/          # Внешние API клиенты
│   ├── model/             # Доменные модели
│   ├── repository/        # Слой доступа к данным
│   │   ├── migrations/    # Миграции базы данных
│   │   └── people/        # Репозиторий людей
│   ├── router/            # HTTP роутер
│   ├── server/            # HTTP сервер
│   └── service/           # Бизнес-логика
├── shared/                # Общие пакеты
│   └── pkg/openapi/       # Сгенерированный OpenAPI код
├── docker-compose.yaml    # Docker Compose конфигурация
├── Dockerfile             # Docker образ
└── Taskfile.yaml          # Задачи для разработки
```

## Быстрый старт

### Предварительные требования

- Go 1.24.3+
- PostgreSQL 15+
- Docker & Docker Compose (опционально)
- Task (опционально, для удобства разработки)

### Запуск с Docker (рекомендуется)

1. **Клонируйте репозиторий**

```bash
git clone https://github.com/forsitet/people-api.git
cd people-api
```

2. **Запустите приложение**

```bash
# С помощью Task
task docker:compose:up

# Или напрямую
docker-compose up -d
```

3. **Откройте Swagger UI**

```
http://localhost:8080/swagger
```

### Запуск локально

1. **Запустите приложение**

```bash
# С помощью Task (рекомендуется)
task docker:compose:up

# Или напрямую
docker-compose up -d
```

2. **Откройте Swagger UI**

```
http://localhost:8080/swagger
```

**Примечание:** Переменные окружения автоматически загружаются из файла `.env` при запуске через Docker Compose.

## API Endpoints

### People

| Метод | Endpoint                  | Описание                                          |
| ---------- | ------------------------- | --------------------------------------------------------- |
| `GET`    | `/api/v1/people`        | Получить список всех людей         |
| `POST`   | `/api/v1/create`        | Создать нового человека              |
| `GET`    | `/api/v1/people/search` | Поиск по ID или фамилии                  |
| `PATCH`  | `/api/v1/people/search` | Обновить информацию о человеке |

### Friends

| Метод | Endpoint                        | Описание                           |
| ---------- | ------------------------------- | ------------------------------------------ |
| `POST`   | `/api/v1/people/{id}/friends` | Добавить друга                |
| `GET`    | `/api/v1/people/{id}/friends` | Получить список друзей |

### Emails

| Метод | Endpoint                       | Описание                  |
| ---------- | ------------------------------ | --------------------------------- |
| `POST`   | `/api/v1/people/{id}/emails` | Добавить email адрес |

## Интеграция с внешними API

Система автоматически получает недостающие данные из следующих сервисов:

- **agify.io** - получение возраста по имени
- **genderize.io** - получение пола по имени
- **nationalize.io** - получение национальности по имени

### Пример использования

```bash
# Создание человека только с именем (данные получаются автоматически)
curl -X POST "http://localhost:8080/api/v1/create" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Иван",
    "surname": "Иванов",
    "patronymic": "Иванович",
    "emails": ["ivan@yandex.ru"]
  }'
```

Система автоматически получит:

- Возраст через agify.io
- Пол через genderize.io
- Национальность через nationalize.io

## Разработка

### Установка зависимостей

```bash
# Обновление зависимостей Go
task deps:update

# Или вручную
go mod download
go mod tidy
```

### Линтинг и форматирование

```bash
# Форматирование кода (автоматически устанавливает инструменты)
task format

# Линтинг (автоматически устанавливает golangci-lint)
task lint
```

### Генерация OpenAPI кода

```bash
# Генерация Go кода из OpenAPI спецификации
task ogen:gen

# Сборка OpenAPI в один файл
task redocly-cli:bundle
```

## Docker команды

### Сборка и запуск

```bash
# Сборка Docker образа
task docker:build

# Запуск собранного образа
task docker:run

# Сборка и запуск с Docker Compose
task docker:compose:up

# Сборка контейнеров с нуля (без кэша)
task docker:compose:build

# Остановка и удаление контейнеров
task docker:compose:down
```

## Конфигурация

### Переменные окружения

| Переменная  | Описание                      | По умолчанию |
| --------------------- | ------------------------------------- | ----------------------- |
| `POSTGRES_HOST`     | Хост PostgreSQL                   | `postgres`            |
| `POSTGRES_PORT`     | Порт PostgreSQL                   | `5432`                |
| `POSTGRES_USER`     | Пользователь БД         | `user`                |
| `POSTGRES_PASSWORD` | Пароль БД                     | `123`                 |
| `POSTGRES_DB`       | Имя базы данных          | `people`              |
| `APP_PORT`          | Порт приложения         | `8080`                |
| `MIGRATIONS_DIR`    | Директория миграций | `migrations`          |

### Миграции

Проект использует [Goose](https://github.com/pressly/goose) для управления миграциями:

- `00001_create_tabels.sql` - создание таблиц
- `00002_seed_data.sql` - заполнение тестовыми данными
- `00003_create_surname_index.sql` - создание индекса по фамилии
