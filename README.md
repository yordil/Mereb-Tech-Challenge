Here's a comprehensive README.md template that outlines the setup, running the application, and accessing Swagger documentation. You can modify it as needed to suit your project specifics.

````markdown
# Mereb Tech Challenge

Welcome to the Mereb Tech Challenge! This project is built using Go, Gin framework, and provides a RESTful API for user management. The application is designed to demonstrate best practices in structuring a Go application and includes features for user creation, retrieval, updating, and deletion.

## Table of Contents

- [Technologies Used](#technologies-used)
- [Folder Structure](#folder-structure)
- [Setup Instructions](#setup-instructions)
- [Running the Application](#running-the-application)
  - [Running Locally](#running-locally)
  - [Running with Docker](#running-with-docker)
- [Swagger Documentation](#swagger-documentation)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Technologies Used

- Go (1.21)
- Gin (web framework)
- Docker
- Swagger for API documentation

## Folder Structure

```plaintext
.
.
├── cmd
│   └── main.go                  # Entry point of the application
├── internal
│   ├── controllers              # Contains controller logic
│   │   └── user_controller.go    # UserController for handling user requests
│   ├── dtos                     # Contains response DTOs
│   │   └── response_dtos.go
│   ├── domain                   # Contains domain models
│   │   └── user.go
│   ├── repository               # Contains data access logic
│   ├── routes                   # Contains route definitions
│   │   ├── user_route.go         # User routes
│   │   └── routes.go             # General route setup
│   └── usecase                  # Contains use case logic
└── docs
```
````

## Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/mereb-tech-challenge.git
   cd mereb-tech-challenge
   ```

2. **Install Go**

   Make sure you have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).

3. **Install Dependencies**

   In the project directory, run:

   ```bash
   go mod tidy
   ```

## Running the Application

### Running Locally

1. **Navigate to the Project Directory**

   Make sure you are in the root directory of the project.

2. **Run the Application**

   Use the following command to start the application:

   ```bash
   go run cmd/main.go
   ```

3. **Access the API**

   Once the application is running, you can access the API at `http://localhost:8080`.

### Running with Docker

1. **Build the Docker Image**

   Make sure Docker is installed and running on your machine. Then, build the Docker image using:

   ```bash
   docker build -t mereb-tech-challenge .
   ```

2. **Run the Docker Container**

   After building the image, you can run it with the following command:

   ```bash
   docker run -p 8080:8080 mereb-tech-challenge
   ```

   Or you can run using compose

   ```bash
    docker-compose up --build

   ```

3. **Access the API**

   Similar to the local run, access the API at `http://localhost:8080`.

## Swagger Documentation

The application includes Swagger documentation to help you understand the available API endpoints.

1. **Generate Swagger Documentation**

   If you modify any endpoints or add new features, regenerate the Swagger documentation by running:

   ```bash
   swag init -g cmd/main.go
   ```

2. **Access Swagger UI**

   Once the application is running, access the Swagger documentation at:

   ```
   http://localhost:8080/swagger/index.html
   ```

## API Endpoints

Here are the available API endpoints:

### User Endpoints

- **Create User**
  - `POST /person`
  - Request Body: `{ "name": "John Doe", "age": 30, "hobbies": ["reading", "swimming"] }`
- **Get All Users**

  - `GET /person`

- **Get User by ID**

  - `GET /person/:personId`

- **Update User**

  - `PUT /person/:personId`
  - Request Body: `{ "name": "Jane Doe", "age": 31, "hobbies": ["reading"] }`

- **Delete User**
  - `DELETE /person/:personId`
