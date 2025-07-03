# Database Schema

This document details the database schema for the Game Leaderboard system, implemented using PostgreSQL.

## Tables

### `users` table
Stores information about registered users.

*   `id` (SERIAL PRIMARY KEY): Unique identifier for the user.
*   `username` (VARCHAR(255) UNIQUE NOT NULL): Unique username of the player.
*   `join_date` (TIMESTAMP DEFAULT CURRENT_TIMESTAMP): Timestamp when the user registered.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### `game_sessions` table
Records individual game sessions and their scores.

*   `id` (SERIAL PRIMARY KEY): Unique identifier for the game session.
*   `user_id` (INT REFERENCES users(id) ON DELETE CASCADE): Foreign key referencing the `users` table, indicating which user played the session.
*   `score` (INT NOT NULL): The score achieved in the game session.
*   `game_mode` (VARCHAR(50) NOT NULL): The mode of the game (e.g., 'solo', 'team').
*   `timestamp` (TIMESTAMP DEFAULT CURRENT_TIMESTAMP): Timestamp when the game session occurred.

```sql
CREATE TABLE game_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    score INT NOT NULL,
    game_mode VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### `leaderboard` table
Stores aggregated total scores and ranks for users, optimized for quick leaderboard retrieval.

*   `id` (SERIAL PRIMARY KEY): Unique identifier for the leaderboard entry.
*   `user_id` (INT REFERENCES users (id) ON DELETE CASCADE UNIQUE): Foreign key referencing the `users` table. `UNIQUE` constraint ensures each user has only one entry.
*   `total_score` (INT NOT NULL): The sum of all scores for the user across all their game sessions.
*   `rank` (INT): The current rank of the user on the leaderboard. This field would typically be updated periodically or on-the-fly.

```sql
CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE UNIQUE,
    total_score INT NOT NULL,
    rank INT
);
```

## Indexes
Indexes are created to improve query performance on frequently accessed columns.

*   `idx_leaderboard_total_score` on `leaderboard(total_score DESC)`: Speeds up queries that order the leaderboard by total score.
*   `idx_leaderboard_user_id` on `leaderboard(user_id)`: Speeds up lookups for a specific user's leaderboard entry.
*   `idx_game_sessions_user_id` on `game_sessions(user_id)`: Speeds up queries to retrieve all game sessions for a particular user.

```sql
CREATE INDEX idx_leaderboard_total_score ON leaderboard(total_score DESC);
CREATE INDEX idx_leaderboard_user_id ON leaderboard(user_id);
CREATE INDEX idx_game_sessions_user_id ON game_sessions(user_id);
```
