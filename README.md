# QR Code Generator API

!['GoLang'](https://golang.org/doc/gopher/fiveyears.jpg)

This project is a simple API for generating QR codes for URLs. It is built using Go and the Fiber framework.

## Prerequisites

Before you can run this project, make sure you have the following installed:

- **Go**: The Go programming language (version 1.16 or higher is recommended). You can download and install Go from [golang.org](https://golang.org/doc/install).
- **Docker**: This project uses Docker for containerization. You can install Docker from [docker.com](https://www.docker.com/get-started).
- **Git**: Version control system to clone the repository. You can install Git from [git-scm.com](https://git-scm.com/downloads).

## Getting Started

Follow these steps to set up and run the project:

### 1. Clone the Repository

First, clone the repository to your local machine using Git:

```bash
git clone git@github.com:Luiggi13/golang.git
cd golang
```


### 2. Set Up the Environment Variables

Create a .env file in the root of the project directory. This file will store environment variables needed to run the project. Here's an example of what your .env file should look like:

```bash
ADDRESS=127.0.0.1
PORT=8080
```


### 3. Install Go Dependencies

Next, install the necessary Go dependencies. Make sure you are in the project directory and run:

```bash
go mod tidy
```

This command will download and install all the required Go packages.

### 4. Run the Project

To run the project, use the following command:

```bash
go run main.go
```

This will start the server on the address and port specified in your .env file (e.g., http://127.0.0.1:8080).

### 5. Testing the API
Once the server is running, you can test the API by sending requests to it. For example, you can use curl or Postman to send a POST request to generate a QR code:

```bash
curl -X POST http://127.0.0.1:8080/api/v1/qr -H "Content-Type: application/json" -d '{"url":"https://example.com"}'
```

> This request should return a base64-encoded QR code image for the provided URL.

## Using Docker (Optional)

If you prefer to run the project in a Docker container, you can do so with the following steps:

### 1. Build the Docker Image

First, build the Docker image:

```bash
docker compose build
```

### 2. Run the Docker Container

Next, run the container:

```bash
docker compose up -d
```

This will start the server in a Docker container, and it will be accessible at http://127.0.0.1:8080 (or the port specified in your .env file).

# Project Structure

- `main.go`: The main entry point for the application.
- `router/`: Contains the route definitions for the API.
- `handler/`: This folder contains the core logic for handling API requests. It includes:
  - `qrcode.go`: Contains the function `CreateQrCode` which handles the creation of QR codes based on the input URL provided in the request.
  - `health.go`: Contains the `GetHealth` function, which is used to check the health status of the API. The `Health` struct is also defined here, representing the structure of the health check response.
- `.env`: Environment variables configuration file (not included). You can clone the `.env.template` file provided in the project and rename it to `.env`. Modify this file according to your environment and project needs, such as changing the server's address and port.

# Troubleshooting

- Error loading .env file: Ensure the .env file is correctly placed in the root directory and contains valid values.
- Port already in use: Make sure the specified port in your .env file is not being used by another application.

# Contributing
> Feel free to fork this repository and submit pull requests. All contributions are welcome!

# License

This project is licensed under the MIT License.
