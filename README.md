Comment Moderation Service

This project implements a RESTful service for comment moderation using Go. The service supports basic CRUD operations for managing comments along with features such as authentication using JWT and middleware for logging and error handling.

Key Features

CRUD Operations: Create, retrieve, update, and delete comments.
Authentication: Secure endpoints using JWT-based authentication.
Middleware: Integrated logging and error handling for request processing.
Containerization: Docker and Docker Compose support for easy deployment.
Modular Architecture: Clean separation of concerns with packages for models, repositories, services, and handlers.
Project Structure

The project follows a layered architecture which includes:

cmd/: The entry point of the application.
internal/: Contains configuration, handlers, middleware, models, repositories, and services.
deployments/: Dockerfile, Docker Compose configuration, and Makefile for building and running the project.
