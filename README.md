# UchetUsers

## 📌 Описание проекта
UchetUsers — это REST API сервис для управления пользователями, построенный на `Golang` с использованием `Gin` и `Swagger`.

## 🚀 Функционал
- Регистрация нового пользователя
- Получение информации о пользователе
- Обновление данных пользователя
- Удаление пользователя
- Логирование запросов
- Документация через Swagger

## 🛠️ Технологии
- Golang
- Gin (web-фреймворк)
- Swagger (документация API)
- Logrus (логирование)
- Docker + Docker Compose
- PostgreSQL (БД)

## 📂 Структура проекта
UchetUsers/ │── internal/ │ ├── handlers/ # Контроллеры (UserHandler) │ ├── models/ # Определение моделей │ ├── services/ # Логика работы с пользователями │ ├── middleware/ # Логирование запросов │── docs/ # Swagger-документация │── main.go # Точка входа │── Dockerfile # Образ для Docker │── docker-compose.yml # Конфигурация Docker Compose │── go.mod # Зависимости │── README.md # Описание проекта

bash
Копировать
Редактировать

## 🔧 Установка и запуск

### 📦 Локальный запуск
1. Установите Go: https://go.dev/dl/
2. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/ITprogDM/User-Accounting.git
   cd UchetUsers
Установите зависимости:

go mod tidy

Сгенерируйте Swagger-документацию:

swag init -g internal/handlers/handler.go -o docs

Запустите сервер:

go run main.go

API доступно по адресу:

http://localhost:8080

Swagger UI:

http://localhost:8080/swagger/index.html

🐳 Запуск в Docker

Соберите и запустите контейнер:

docker-compose up --build

API доступно на http://localhost:8080

Swagger-документация:

http://localhost:8080/swagger/index.html

📄 API Методы

🟢 Создание пользователя
POST /users

json
{
"name": "Иван Иванов",
"email": "ivan@example.com",
"age": 30
}
Ответ:

json
{
"message": "Пользователь успешно создан"
}
🔵 Получение пользователя
GET /users/{id}
Ответ:

json
{
"id": 1,
"name": "Иван Иванов",
"email": "ivan@example.com",
"age": 30
}
🟡 Обновление пользователя
PUT /users/{id}

json
{
"id": 1,
"name": "Иван Петров",
"email": "ivanp@example.com",
"age": 31
}
Ответ:

json
{
"message": "Пользователь успешно обновлён"
}
🔴 Удаление пользователя
DELETE /users/{id}
Ответ:

json
{
"message": "Пользователь успешно удалён"
}
