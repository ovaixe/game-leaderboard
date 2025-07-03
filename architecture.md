# System Architecture

This document outlines the high-level and low-level design of the gaming leaderboard system.

## High-Level Design (HLD)

The system will consist of the following components:

*   **Go Backend:** A web server that exposes a RESTful API for submitting scores and retrieving leaderboard data.
*   **Postgres Database:** A relational database that stores user data, game sessions, and leaderboard rankings.
*   **Python Simulation Script:** A script that simulates user activity by sending requests to the API.

## Low-Level Design (LLD)

### Go Backend

*   **`main.go`:** The entry point of the application. It will initialize the database connection, set up the HTTP router, and start the web server.
*   **`handlers.go`:** This file will contain the HTTP handlers for the API endpoints.
*   **`storage.go`:** This file will handle all interactions with the Postgres database.
*   **`models.go`:** This file will define the data structures for users, game sessions, and leaderboard entries.

### Database Schema

*   **`users` table:**
    *   `id` (SERIAL PRIMARY KEY)
    *   `username` (VARCHAR(255) UNIQUE NOT NULL)
    *   `join_date` (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)
*   **`game_sessions` table:**
    *   `id` (SERIAL PRIMARY KEY)
    *   `user_id` (INT REFERENCES users(id) ON DELETE CASCADE)
    *   `score` (INT NOT NULL)
    *   `game_mode` (VARCHAR(50) NOT NULL)
    *   `timestamp` (TIMESTAMP DEFAULT CURRENT_TIMESTAMP)
*   **`leaderboard` table:**
    *   `id` (SERIAL PRIMARY KEY)
    *   `user_id` (INT REFERENCES users(id) ON DELETE CASCADE)
    *   `total_score` (INT NOT NULL)
    *   `rank` (INT)
