# MailQueue

A simple asynchronous email queue system built in Go using RabbitMQ for reliable message processing.

## Features

- **Asynchronous Email Sending:** Emails are placed in a queue and processed asynchronously, preventing blocking of the main application thread.
- **Reliable Delivery:** RabbitMQ ensures that emails are not lost in case of application failure.
- **Scalable:** The worker system can be scaled independently to handle a high volume of emails.
- **Simple API:** A straightforward API to enqueue emails.

## Project Structure

```
/
├── cmd/
│   └── server/
│       └── main.go         # Main application entry point
├── handlers/
│   ├── email.go            # HTTP handlers for email-related requests
│   └── health.go           # HTTP handler for health checks
├── internal/
│   ├── mail/
│   │   └── send.go         # Logic for sending emails
│   ├── models/
│   │   └── email.go        # Email data structure
│   └── rabbitmq/
│       ├── connection.go   # RabbitMQ connection management
│       ├── consumer.go     # RabbitMQ consumer logic
│       └── producer.go     # RabbitMQ producer logic
├── routes/
│   └── router.go           # API route definitions
├── .env                    # Environment variables
├── .gitignore              # Git ignore file
├── docker-compose.yml      # Docker Compose configuration
├── Dockerfile              # Dockerfile for building the application
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── README.md               # This file
```

## Requirements

- Go 1.24 or higher
- Docker

## Getting Started

### With Docker (Recommended)

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/mailqueue.git
   cd mailqueue
   ```

2. **Create a `.env` file:**

   Copy the example environment variables:

   ```bash
   cp .env.example .env
   ```

   Update the `.env` file with your configuration.

3. **Build and run with Docker Compose:**

   ```bash
   docker-compose up --build
   ```

   The application will be available at `http://localhost:8080`.

### Without Docker

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/mailqueue.git
   cd mailqueue
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**

   Create a `.env` file in the root of the project and add the following:

   ```
   RABBITMQ_URL=amqp://guest:guest@localhost:5672/
   PORT=8080
   ```

4. **Run RabbitMQ:**

   If you have Docker installed, you can easily start a RabbitMQ instance:

   ```bash
   docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
   ```

5. **Run the application:**

   ```bash
   go run cmd/server/main.go
   ```

   The server will start on port 8080.

## API Endpoints

### Enqueue Email

- **URL:** `/api/email`
- **Method:** `POST`
- **Body:**

  ```json
  {
    "to": "recipient@example.com",
    "subject": "Hello from MailQueue",
    "body": "This is a test email."
  }
  ```

- **Success Response:**

  - **Code:** `202 Accepted`
  - **Content:**

    ```json
    {
      "message": "Email enqueued"
    }
    ```

### Health Check

- **URL:** `/api/health`
- **Method:** `GET`
- **Success Response:**

  - **Code:** `200 OK`
  - **Content:**

    ```json
    {
      "status": "ok"
    }
    ```

## Dependencies

- [chi v5.2.2](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router for building Go HTTP services.
- [godotenv v1.5.1](https://github.com/joho/godotenv) - A Go port of the Ruby dotenv library (Loads environment variables from `.env`).
- [amqp v1.1.0](https://github.com/streadway/amqp) - Go client for RabbitMQ.

## To-Do

- [ ] Add unit and integration tests
- [ ] Implement a more robust logging solution
- [ ] Add support for attachments in emails
- [ ] Create a more detailed documentation for the API

## License

This project is licensed under the MIT License.