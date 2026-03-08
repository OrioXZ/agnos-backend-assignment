# 🏥 Agnos Backend Assignment

Backend service built with **Go (Gin)** implementing hospital-scoped
authentication and patient search system.

------------------------------------------------------------------------

## 🚀 Tech Stack

-   Go
-   Gin Framework
-   PostgreSQL
-   Docker Compose
-   Nginx (Reverse Proxy)

------------------------------------------------------------------------

## 📌 Overview

This system implements:

-   Staff registration & login
-   JWT-based authentication
-   Hospital-scoped access control
-   Patient search APIs
-   Containerized deployment using Docker Compose

Each staff member belongs to a hospital and can only access patients
within their own hospital.

------------------------------------------------------------------------

# 🧱 Architecture

Handler → Service → Repository → Database

-   **Handler**: HTTP layer & DTO binding\
-   **Service**: Business logic & validation\
-   **Repository**: GORM database operations\
-   **Middleware**: JWT authentication & hospital isolation

------------------------------------------------------------------------

# 📡 API Endpoints

## 1️⃣ Create Staff

POST `/staff/create`

## 2️⃣ Login

POST `/staff/login`

## 3️⃣ Search Patient by ID

GET `/patient/search/{id}`

## 4️⃣ Search Patient (Filter)

GET `/patient/search`

Optional query parameters: - first_name - last_name - national_id -
passport_id - date_of_birth - phone_number - email

------------------------------------------------------------------------

# 🐳 Running with Docker

## Build & Run

    docker compose up --build

Access:

    http://localhost

Health Check:

    GET /health

------------------------------------------------------------------------

# 🧪 Testing

Run tests:

    go test ./...

------------------------------------------------------------------------

# 🔒 Security

-   Password hashing using bcrypt
-   JWT-based authentication
-   Hospital-level data isolation via middleware

------------------------------------------------------------------------

# 📦 Deliverables

-   GitHub repository
-   Google Doc documentation
-   Docker Compose setup
-   Docker Compose configuration is included for easier local setup.

------------------------------------------------------------------------

## 👨‍💻 Author

Backend assignment implemented using Go & Gin with layered architecture
and secure hospital data isolation.
