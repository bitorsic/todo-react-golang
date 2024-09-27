
# Task-inator 3000

Task-inator 3000 is a simple and efficient task management application that allows users to create and manage their task lists with ease. The application is built using a modern tech stack to ensure a seamless user experience, both on the frontend and backend.

## Features

- **Task Management**: Create, read, update, and delete tasks within task lists.
- **Task Lists**: Organize tasks into multiple lists for better productivity.
- **JWT Authentication**: Secure access to the application with JWT-based authentication.
- **Dark Mode**: Full support for dark mode across the application.
- **Responsive Design**: Accessible across all screen sizes and devices.

## Tech Stack

- **Frontend**: React, TypeScript, Tailwind CSS
- **Backend**: Go Fiber, MongoDB, Redis
- **Authentication**: JWT (JSON Web Tokens)

## Installation

### Prerequisites

- Node.js (v16+)
- Go (v1.18+)
- MongoDB (local or cloud)
- Redis (for token management and task queues)

### Setup Instructions

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/bitorsic/todo-react-golang.git
   cd todo-react-golang
   ```

2. **Backend Setup (Go Fiber):**

   - Navigate to the backend directory:
   
     ```bash
     cd server
     ```
   
   - Install Go dependencies:

     ```bash
     go mod download
     ```

   - Set up environment variables (`.env` file):
   
     ```bash
     MONGODB_URI="mongodb://localhost:27017/taskinator"
     REDIS_URL="redis://localhost:6379"
     AUTH_TOKEN_KEY="some_key_you_can_use"
     REFRESH_TOKEN_KEY="another_key"
     FRONTEND_URL="http://localhost:5173"
     ```

   - Run the backend:

     ```bash
     go run main.go
     ```

3. **Frontend Setup (React TypeScript):**

   - Navigate to the frontend directory:
   
     ```bash
     cd client
     ```
   
   - Install frontend dependencies:

     ```bash
     npm install
     ```

   - Create a `.env` file and specify the backend URL:

     ```bash
     VITE_BACKEND_URL="http://localhost:3000"
     ```

   - Run the React development server:

     ```bash
     npm run dev
     ```

4. **Redis Setup:**

   - Make sure Redis is installed and running locally or in the cloud.
   - Task-inator 3000 uses Redis for token blacklisting and managing task queues.

## Usage

### Creating Task Lists

- After logging in, you can create a new task list from the dashboard.
- Each task list can have multiple tasks, which can be added, updated, or deleted.

### Task Management

- Add a task to a task list via the "Add Task" button.
- Tasks can be marked as completed or deleted from the list.

### Authentication

- Task-inator 3000 uses JWT-based authentication. When a user logs in, two JWT tokens are issued:
  - **Access Token**: Expires in 10 minutes.
  - **Refresh Token**: Expires in 30 days, used to get a new access token when the original expires.

## Deployment

### Backend Deployment

1. **Build the Go application**:

   ```bash
   go build -o taskinator-backend
   ```

2. **Deploy** the built application to your preferred cloud provider.

### Frontend Deployment

1. **Build the frontend**:

   ```bash
   npm run build
   ```

2. Serve the static files in /dist directory using any web server (e.g., NGINX).

## License

Task-inator 3000 is open-source software licensed under the [MIT license](LICENSE).
