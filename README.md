# MyChat

![Go Version](https://img.shields.io/badge/Go-1.XX-blue)

MyChat — это веб-приложение для общения в реальном времени, написанное на Golang.

## 🚀 Возможности
- Регистрация и авторизация пользователей
- Отправка и получение сообщений в реальном времени
- Групповые и личные чаты
- WebSocket для моментальной передачи сообщений
- Хранение сообщений в базе данных

## 🛠️ Технологии
- **Backend:** Golang (Fiber, WebSocket)
- **Database:** PostgreSQL
- **Auth:** JWT

## 📦 Установка и запуск
### 1. Клонирование репозитория
```sh
git clone https://github.com/terfo1/MyChat.git
cd MyChat
```

### 2. Настройка переменных окружения
Создайте файл `.env` и укажите настройки, например:
```ini
DB="host=yourhost user=youruser password=yourpassword dbname=mychat port=5432 sslmode=disable"
SECRET_KEY=your_jwtsecret_key
```

### 3. Установка зависимостей
```sh
go mod tidy
```

### 4. Запуск сервера
```sh
go run main.go
```

## 🔗 API Роуты
| Метод  | Роут                | Описание                                    |
|--------|---------------------|---------------------------------------------|
| GET    | `/api/`             | Главная страница                            |
| POST   | `/api/signup`       | Регистрация пользователя                    |
| POST   | `/api/login`        | Авторизация пользователя                    |
| GET    | `/api/profile`      | Получение профиля (требуется аутентификация) |
| GET    | `/api/ws/:id`       | Подключение к WebSocket-чату (аутентификация) |
| GET    | `/api/messages/:user1/:user2` | Получение истории чатов между пользователями |

## 👥 Контакты
Если у вас есть вопросы или предложения, создайте issue или свяжитесь со мной:
- GitHub: [terfo1](https://github.com/terfo1)

