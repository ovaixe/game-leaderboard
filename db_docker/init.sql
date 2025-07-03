CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    score INT NOT NULL,
    game_mode VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE UNIQUE,
    total_score INT NOT NULL,
    rank INT
);

-- Indexes for performance
CREATE INDEX idx_leaderboard_total_score ON leaderboard(total_score DESC);
CREATE INDEX idx_leaderboard_user_id ON leaderboard(user_id);
CREATE INDEX idx_game_sessions_user_id ON game_sessions(user_id);

-- Populate Users Table with 1 Million Records
 INSERT INTO users (username)
 SELECT 'user' || generate_series(1, 1000000);

-- Populate Game Sessions with Random Scores
 INSERT INTO game_sessions (user_id, score, game_mode, timestamp)
 SELECT
     floor(random() * 1000000 + 1)::int,
     floor(random() * 10000 + 1)::int,
     CASE WHEN random() > 0.5 THEN 'solo' ELSE 'team' END,
     NOW() - INTERVAL '1 day' * floor(random() * 365)
 FROM generate_series(1, 5000000);

-- Populate Leaderboard by Aggregating Scores
 INSERT INTO leaderboard (user_id, total_score, rank)
 SELECT user_id, AVG(score) as total_score, RANK() OVER (ORDER BY SUM(score) DESC)
 FROM game_sessions
 GROUP BY user_id;
