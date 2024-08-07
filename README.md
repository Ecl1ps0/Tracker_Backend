
# Proctor API

This project is a server service for the Proctor application, created using the Gin web-framework. It includes connection via WebSocket, data caching using Redis, connection to a PostgreSQL database, Docker containerization, API for managing users, tasks, solutions and reports, as well as a machine learning model for determining the percentage of student cheating.
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_PASSWORD`

`HASH_SALT`

`SIGN_KEY`


## Installation and Launch

1. Clone the repository:
```bash
    git clone https://github.com/Ecl1ps0/Tracker_Backend.git
    cd Tracker_Backend
```

2. Build and start the containers:
```bash
docker-compose up --build
```



## API Reference

### Authentication

- **POST** `/auth/signin`  
  Sign in a user

- **POST** `/auth/signup`  
  Sign up a new user

### WebSocket

- **GET** `/connection/ws`  
  Establish a WebSocket connection between Backend and Keylogger

### Users

- **GET** `/api/users`  
  Get all users

- **GET** `/api/users/profile`  
  Get the profile of the authenticated user

- **GET** `/api/users/students`  
  Get all students

- **GET** `/api/users/on-solution/:id`  
  Get a student by solution ID

- **GET** `/api/users/teacher/:id/students`  
  Get students of a specific teacher

- **POST** `/api/users/:studentID/add-to-task/:taskID`  
  Add a student to a task

### Tasks

- **GET** `/api/tasks`  
  Get all tasks

- **GET** `/api/tasks/:id`  
  Get a task by ID

- **GET** `/api/tasks/teacher/:id`  
  Get all tasks of a specific teacher

- **GET** `/api/tasks/student/:id`  
  Get all tasks of a specific student

- **POST** `/api/tasks/create`  
  Create a new task

- **PUT** `/api/tasks/update/:id`  
  Update a task

- **DELETE** `/api/tasks/delete/:id`  
  Delete a task

### Solutions

- **GET** `/api/solutions`  
  Get all solutions

- **GET** `/api/solutions/by-student/:id`  
  Get solutions by student ID

- **GET** `/api/solutions/solved-task/:id`  
  Get student solutions for a specific task

- **GET** `/api/solutions/get-solution-on-student-task/:id`  
  Get a solution for a student's task

- **POST** `/api/solutions/on-student-task/:id`  
  Create a solution for a student's task

- **PUT** `/api/solutions/generate-cheating-rate/:id`  
  Generate a cheating rate for a solution

- **PUT** `/api/solutions/update-final-grade/:id`  
  Update the final grade of a solution

### Reports

- **GET** `/api/reports`  
  Get all reports

- **POST** `/api/reports/createReport`  
  Create a new report

### Swagger

Swagger documentation is available on `http://localhost:8080/swagger/index.html`

