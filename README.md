# Todo API

This project is a RESTful ToDo API written in Go using the Gin framework.

## Features

- Create, read, update, and delete tasks (CRUD).
- Health check endpoint to verify server status.
- Middleware for logging requests.

## Technologies Used

- **Golang**: Main programming language.
- **Gin**: HTTP framework for RESTful APIs.
- **Postman**: API testing.
- **Swaggo**: Auto-generated API documentation with Swagger (optional).

## Installation

### Prerequisites

- Install [Go](https://golang.org/doc/install).
- Set up [Git](https://git-scm.com/).
- Optional: Install [Postman](https://www.postman.com/) to test the endpoints.

### Installation Steps

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/todo-api.git

2. Navigate to the project directory:
    ```bash
   cd todo-api

3. Install dependencies:
    ```bash
    go mod tidy

4. Run the server
    ```bash
    go run cmd/main.go

## Endpoints

### Health Chech

Verify that the server is running.

    - GET /health
        {
            "status": "ok"
        }
ToDo Tasks
    - GET /api/v1/todos
    Retrieves all tasks.

    - POST /api/v1/todos
    Creates a new task.

    **Requested Body (JSON)**
    {
        "title": "Task title",
        "description": "Task description"
    }

    - GET /api/v1/todos/:id
    Retrieves a specific task by its ID.

    - PUT /api/v1/todos/:id
    Updates a existing task.

    Request Body (JSON):
    {
        "title": "Updated title",
        "description": "Updated description"
    }

    - DELETE /api/v1/todos/:id
    Deletes a specific task by its ID.

## Project Structure

/todo-api
├── cmd/
│   └── main.go             # Main program entry point
├── internal/
│   ├── handlers/           # Endpoint logic
│   │   └── todo_handler.go
│   ├── models/             # Data models
│   │   └── todo.go
│   ├── repositories/       # Database interaction logic
│   │   └── todo_repo.go
│   ├── routes/             # Route configuration
│   │   └── routes.go
│   └── services/           # Business logic
│       └── todo_service.go
├── configs/                # Configuration files (e.g., JSON, env vars)
├── docs/                   # Documentation (Swagger or others)
├── go.mod                  # Project dependencies
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation

## API Documentation (Swagger)
This project uses Swag to auto-generate API documentation.

1. Install Swag:
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest

2. Generate the documentation:
    ```bash
    swag init

3. The documentation will be available at /swagger/index.html if the docs folder is integrated into the server

## Testing with Postman
You can import the following endpoints into Postman to test the API:

1. Create a new collection.

2. Add the following requests:
    *GET* /health
    *GET* /api/v1/todos
    *POST* /api/v1/todos with a JSON body.
    *GET* /api/v1/todos/:id
    *PUT* /api/v1/todos/:id with a JSON body.
    *DELETE* /api/v1/todos/:id

3. Execute the requests one by one to verify functionality.

## Contributions
Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository
2. Create a branch for you feature or fix:
    ```bash
    git checkout -b feature/new-feature

3. Make the necessary changes and commit them:
    ```bash
    git commit -m "Add new feature"

4. Push your changes and submit a pull request.

## License
This project is licensed under the MIT License.


Thank you for checking out this project!