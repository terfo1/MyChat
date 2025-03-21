# MyChat

![Go Version](https://img.shields.io/badge/Go-1.23-blue)

MyChat is a real-time web chat application built with Golang.

## 🚀 Features
- User registration and authentication
- Real-time messaging
- Private chats
- WebSocket for instant message delivery
- Message storage in a database

## 🛠️ Technologies
- **Backend:** Golang (Fiber, WebSocket)
- **Database:** PostgreSQL
- **Auth:** JWT

## 📦 Installation and Setup
### 1. Clone the repository
```sh
git clone https://github.com/terfo1/MyChat.git
cd MyChat
```

### 2. Configure environment variables
Create a `.env` file and specify your settings, e.g.:
```ini
DB="host=yourhost user=youruser password=yourpassword dbname=mychat port=5432 sslmode=disable"  
SECRET_KEY=your_jwt_secret_key
```

### 3. Install dependencies
```sh
go mod tidy
```

### 4. Run the server
```sh
go run main.go
```

## 🔗 API Routes
| Method  | Route                     | Description                                    |
|---------|---------------------------|----------------------------------------------|
| GET     | `/api/`                    | Home page                                    |
| POST    | `/api/signup`              | User registration                            |
| POST    | `/api/login`               | User authentication                         |
| GET     | `/api/profile`             | Get user profile (authentication required)  |
| GET     | `/api/ws/:id`              | Connect to WebSocket chat (authentication required) |
| GET     | `/api/messages/:user1/:user2` | Retrieve chat history between users       |


## 👥 Contact
If you have any questions or suggestions, feel free to create an issue or contact me:
- GitHub: [terfo1](https://github.com/terfo1)
