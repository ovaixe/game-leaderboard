# API Endpoints

This document outlines the RESTful API endpoints exposed by the Go backend.

## Base URL
`/api`

## Endpoints

### 1. Submit Score
*   **Method:** `POST`
*   **Path:** `/scores/submit`
*   **Description:** Submits a new game score for a user.
*   **Request Body:**
    ```json
    {
        "user_id": 123,
        "score": 1000,
        "game_mode": "survival"
    }
    ```
*   **Response (Success - 201 Created):**
    ```json
    {
        "message": "Score submitted successfully",
        "session_id": 456
    }
    ```
*   **Response (Error - 400 Bad Request):**
    ```json
    {
        "error": "Invalid user ID or score"
    }
    ```

### 2. Get Leaderboard
*   **Method:** `GET`
*   **Path:** `/leaderboard`
*   **Description:** Retrieves the global leaderboard.
*   **Query Parameters:**
    *   `limit` (optional, int): Maximum number of entries to return (default: 10).
    *   `offset` (optional, int): Number of entries to skip (default: 0).
*   **Response (Success - 200 OK):**
    ```json
    [
        {
            "user_id": 1,
            "username": "player1",
            "total_score": 5000,
            "rank": 1
        },
        {
            "user_id": 2,
            "username": "player2",
            "total_score": 4500,
            "rank": 2
        }
    ]
    ```

### 3. Get User Rank
*   **Method:** `GET`
*   **Path:** `/leaderboard/user/{user_id}`
*   **Description:** Retrieves the rank and total score for a specific user.
*   **Path Parameters:**
    *   `user_id` (int, required): The ID of the user.
*   **Response (Success - 200 OK):**
    ```json
    {
        "user_id": 123,
        "username": "playerX",
        "total_score": 3200,
        "rank": 5
    }
    ```
*   **Response (Error - 404 Not Found):**
    ```json
    {
        "error": "User not found or not on leaderboard"
    }
    ```
