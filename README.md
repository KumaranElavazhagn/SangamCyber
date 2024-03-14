## Sangam Cyber

sangamCyber is a Go project aimed at providing user authentication and user information management through HTTP APIs. It utilizes the Gorilla Mux router and CORS middleware for handling HTTP requests. The main goal of the project is to securely store passwords in the database using the Tamil language, providing enhanced security and user experience.

### Project Structure

The project is structured as follows:

- **converter**: Contains utility functions for converting between English and Tamil characters.
- **database**: Handles database initialization and configuration.
- **dto**: Defines data transfer objects used in the project.
- **entity**: Defines entity structures representing database entities.
- **errs**: Contains error handling utilities and error response structures.
- **handler**: Implements HTTP request handlers for different endpoints.
- **repository**: Implements repository interfaces for database operations.
- **service**: Implements business logic for handling user information and authentication.

### Installation and Setup

1. **Clone the repository:**
```bash
git clone <[repository_url](https://github.com/KumaranElavazhagn/SangamCyber.git)>
```
2. **Navigate to the project directory:**
```bash
cd sangamCyber
```
3. **Install dependencies:
```bash
go mod tidy
```
4. **Build the project:
```bash
go build
```
5. **Run the project:
```bash
go run main.go
```

By default, the project listens on port 8080.

### Endpoints
- POST /user/info: Insert user information.
- POST /user/auth: Authenticate user information.