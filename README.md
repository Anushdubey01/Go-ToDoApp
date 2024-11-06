# Go TodoApp

A simple, fast, and efficient RESTful API built with Go for managing a to-do list. This application allows users to perform CRUD (Create, Read, Update, Delete) operations on to-do tasks. The API is designed with scalability and ease of use in mind, making it a perfect beginner project for learning Go and RESTful API development.

## Features

- **Create** a new task
- **Read** tasks (all tasks or individual tasks by ID)
- **Update** task details (mark as completed, change title)
- **Delete** a task from the list
- Simple in-memory data storage for persistence (can be easily swapped with a database)

## Table of Contents

- [Technologies Used](#technologies-used)
- [Installation Instructions](#installation-instructions)
- [API Endpoints](#api-endpoints)
- [Running the Application](#running-the-application)
- [Deployment](#deployment)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Technologies Used

This project is built with the following technologies:

- **Go (Golang)** - The main programming language used to create the API.
- **net/http** - The Go standard library for HTTP server and routing.
- **JSON** - Data exchange format for API communication.
- **In-memory storage** - Tasks are stored in memory as a simple key-value store. Can be upgraded to a database later.

## Installation Instructions

### Prerequisites

1. **Go (version 1.18 or higher)** - You need to have Go installed on your local machine. Follow the official instructions to install Go: [Go Installation](https://golang.org/doc/install).
2. **Git** - Required to clone the repository.

### Steps to Setup Locally

1. **Clone the Repository**

   First, clone this repository to your local machine:

   ```bash
   git clone https://github.com/Anushdubey01/Go-TodoApp.git
   cd Go-TodoApp
   ```

2. **Install Dependencies**

   This project uses Go modules, so the dependencies are already listed in `go.mod`. You can download and install them by running:

   ```bash
   go mod tidy
   ```

3. **Run the Application**

   After setting up, run the application with:

   ```bash
   go run cmd/server/main.go
   ```

   By default, the application will run on [http://localhost:8080](http://localhost:8080).

### Running Tests

If you'd like to test the application, you can write tests using Go's built-in testing framework. For example, you can create a `main_test.go` file and run tests using the following command:

```bash
go test ./...
```

## API Endpoints

The Go TodoApp exposes the following RESTful API endpoints:

### **1. GET /tasks**

- **Description**: Fetch all tasks in the to-do list.
- **Response**: A list of all tasks in JSON format.

#### Example Request:

```bash
GET http://localhost:8080/tasks
```

#### Example Response:

```json
[
  {
    "id": 1,
    "title": "Finish Go project",
    "completed": false
  },
  {
    "id": 2,
    "title": "Learn REST APIs",
    "completed": true
  }
]
```

### **2. GET /tasks/{id}**

- **Description**: Fetch a specific task by its ID.
- **Response**: The task with the given ID in JSON format.

#### Example Request:

```bash
GET http://localhost:8080/tasks/1
```

#### Example Response:

```json
{
  "id": 1,
  "title": "Finish Go project",
  "completed": false
}
```

### **3. POST /tasks/create**

- **Description**: Create a new task.
- **Request Body**: A JSON object with `title` (task title) and `completed` (task completion status).
- **Response**: The created task with an ID assigned.

#### Example Request:

```bash
POST http://localhost:8080/tasks/create
Content-Type: application/json

{
  "title": "Buy groceries",
  "completed": false
}
```

#### Example Response:

```json
{
  "id": 3,
  "title": "Buy groceries",
  "completed": false
}
```

### **4. PUT /tasks/update/{id}**

- **Description**: Update a task by its ID. You can modify the title and completion status of the task.
- **Request Body**: A JSON object with the updated `title` and `completed` status.
- **Response**: The updated task.

#### Example Request:

```bash
PUT http://localhost:8080/tasks/update/1
Content-Type: application/json

{
  "title": "Finish Go project - Updated",
  "completed": true
}
```

#### Example Response:

```json
{
  "id": 1,
  "title": "Finish Go project - Updated",
  "completed": true
}
```

### **5. DELETE /tasks/delete/{id}**

- **Description**: Delete a task by its ID.
- **Response**: HTTP status code 204 (No Content) indicating the task was deleted successfully.

#### Example Request:

```bash
DELETE http://localhost:8080/tasks/delete/1
```

#### Example Response:

```plaintext
HTTP/1.1 204 No Content
```

## Running the Application

To run the application locally, follow the instructions in the **Installation Instructions** section.

1. Clone the repository.
2. Install dependencies using `go mod tidy`.
3. Start the server using `go run cmd/server/main.go`.
4. The server will be available at `http://localhost:8080`.

## Deployment

### Deploying on Heroku

To deploy this application on Heroku, follow the steps below:

1. Create an account on [Heroku](https://signup.heroku.com/).
2. Install the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli).
3. From the root of your project, create a Heroku app:

   ```bash
   heroku create
   ```

4. Push your code to Heroku:

   ```bash
   git push heroku main
   ```

5. Once the deployment is complete, you can access your app at:

   ```
   https://your-app-name.herokuapp.com
   ```

### Environment Variables (optional)

Currently, the app uses in-memory storage, so no additional environment variables are needed. If you plan to connect the app to a database like PostgreSQL, you'll need to set the database URL as an environment variable.

## Testing

While the application doesn’t have specific automated tests yet, you can manually test the API endpoints using tools like **Postman** or **cURL**.

For automated testing, you can write test cases using Go's built-in testing framework.

## Contributing

Contributions to this project are welcome! If you have an idea for improving the app or fixing a bug, feel free to submit a pull request.

### Steps for Contributing:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.
