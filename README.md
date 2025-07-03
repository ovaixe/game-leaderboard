# Game Leaderboard System

This repository contains a full-stack game leaderboard system, featuring a Go backend, PostgreSQL database, Redis cache, and a React frontend. It also includes a Python script for simulating game activity.

## Table of Contents
- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features
*   **User Registration:** Register new players.
*   **Score Submission:** Submit game scores for registered users.
*   **Leaderboard Retrieval:** Get global leaderboard rankings.
*   **User Rank Lookup:** Check a specific user's rank and total score.
*   **Scalable Backend:** High-performance Go API.
*   **Efficient Caching:** Redis for fast leaderboard lookups.
*   **Modern Frontend:** Responsive React UI.
*   **Dockerized Deployment:** Easy setup and consistent environments with Docker Compose.
*   **Simulation Script:** Python script to simulate user activity and load.

## Architecture
For a detailed overview of the system's architecture, please refer to [architecture.md](architecture.md) and [system-design.md](system-design.md).

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
*   [Docker](https://www.docker.com/get-started) (Docker Engine and Docker Compose)

### Installation
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/game-leaderboard.git
    cd game-leaderboard
    ```

2.  **Build and run the Docker containers:**
    ```bash
    docker-compose up --build
    ```
    This command will:
    *   Build the Go backend and React frontend Docker images.
    *   Start the PostgreSQL database container (initialized with data from `db_docker/init.sql`).
    *   Start the Redis cache container.
    *   Start the Go backend API server.
    *   Start the React development server.

    It might take a few minutes for all services to start up and for the database to be populated.

### Running the Application
*   **Backend API:** The Go backend will be accessible at `http://localhost:8080`.
*   **Frontend UI:** The React frontend will be accessible at `http://localhost:5173` (or another port if 5173 is in use).
*   **Python Simulation:** To run the simulation script (after all services are up):
    ```bash
    python simulate.py
    ```
    (You might need to install Python dependencies: `pip install requests`)

## API Endpoints
Detailed API documentation can be found in [api-endpoints.md](api-endpoints.md).

## Project Structure
```
. (root)
├── architecture.md             # High-level system architecture
├── api-endpoints.md            # Detailed API documentation
├── database-schema.md          # Database schema details
├── docker-compose.yml          # Docker Compose configuration
├── plan.md                     # Project development plan
├── requirements.md             # Project requirements
├── simulate.py                 # Python script for simulating user activity
├── system-design.md            # Detailed system design document
├── backend/                    # Go backend service
│   ├── cmd/server/main.go      # Main application entry point
│   ├── internal/               # Internal packages (config, controllers, models, etc.)
│   └── pkg/                    # External packages (database, redis clients)
├── db_docker/                  # PostgreSQL Docker setup
│   └── init.sql                # Database initialization script
└── frontend/                   # React frontend application
    ├── src/                    # React source code
    └── ...
```

## Contributing
Contributions are welcome! Please see `CONTRIBUTING.md` (if available) for details on how to contribute.

## License
This project is licensed under the MIT License - see the `LICENSE` file for details.
