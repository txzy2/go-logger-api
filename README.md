# Logger Go API

Go REST API приложение для системы логирования инцидентов, построенное по принципам Clean Architecture.

## Технологический стек

- **Go 1.25.1** - основной язык программирования
- **Gin** - веб-фреймворк для HTTP API
- **GORM** - ORM для работы с PostgreSQL
- **PostgreSQL** - база данных
- **Docker** - контейнеризация

## Архитектура приложения

Приложение построено по принципам Clean Architecture с четким разделением на слои:

```
app/
├── cmd/app/                    # Точка входа приложения
│   └── main.go                 # Главный файл приложения
├── cmd/migrate/                # Миграции базы данных
│   └── main.go                 # Скрипт миграций
├── config/                     # Конфигурация приложения
│   └── app.go                  # Инициализация и запуск сервера
├── internal/                   # Внутренняя логика приложения
│   ├── delivery/http/v1/       # HTTP API handlers
│   ├── service/                # Бизнес-логика
│   ├── repository/             # Слой доступа к данным
│   └── models/                 # Модели данных
├── pkg/                        # Переиспользуемые пакеты
│   ├── database/               # Конфигурация БД
│   └── env.go                  # Загрузка переменных окружения
└── docs/                       # Swagger документация
```

## Модели данных

### Incident (Инцидент)
Основная модель для хранения информации об инцидентах:
- ID, Service, Level, Message
- IncidentTypeID (связь с типом инцидента)
- Action, AdditionalFields, Function, Class, File
- Date, Count
- Timestamps (CreatedAt, UpdatedAt, DeletedAt)

### IncidentType (Тип инцидента)
Классификация типов инцидентов

### Services (Сервисы)
Управление сервисами системы:
- Name, Active status
- Timestamps

### SendTemplate (Шаблоны отправки)
Шаблоны для отправки уведомлений

## Конфигурация

### Переменные окружения

Приложение использует следующие переменные окружения:

```env
# База данных
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASS=your_password
DB_NAME=your_database
DB_SSLMODE=disable

# Сервер
SERVER_PORT=8080
GIN_MODE=release  # или debug для разработки
```

## Сборка и запуск

### Локальная разработка

```bash
# Установка зависимостей
go mod download

# Запуск приложения
go run cmd/app/main.go

# Запуск миграций
go run cmd/migrate/main.go -action=migrate
```

### Docker (через docker-compose)

```bash
# Запуск всех сервисов
docker-compose -f docker-compose.dev.yml up -d

# Перезапуск с пересборкой
make dbr-dev

# Просмотр логов
make dlogs-dev
```

## API Endpoints

### Базовые маршруты

- `GET /api/v1/health` - Проверка состояния приложения
- `GET /api/v1/ping` - Проверка подключения к базе данных

## Основные зависимости

- **gin-gonic/gin** - HTTP веб-фреймворк
- **gorm.io/gorm** - ORM для работы с PostgreSQL
- **gorm.io/driver/postgres** - PostgreSQL драйвер для GORM
- **joho/godotenv** - Загрузка переменных окружения

## Особенности

- **Graceful Shutdown** - корректное завершение работы с обработкой сигналов
- **GORM** - автоматические миграции и управление схемой БД
- **Clean Architecture** - четкое разделение слоев приложения
- **Docker** - полная контейнеризация для разработки и продакшена
