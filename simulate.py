import requests
import random
import time

API_BASE_URL = "http://localhost:8080/api/leaderboard"

# Simulate score submission
def submit_score(user_id):
    score = random.randint(100, 10000)
    start_time = time.time()
    requests.post(f"{API_BASE_URL}/submit", json={"user_id": user_id, "score": score})
    end_time = time.time()
    print(f"SubmitScore API call took: {end_time - start_time:.4f} seconds")

# Fetch top players
def get_top_players():
    start_time = time.time()
    response = requests.get(f"{API_BASE_URL}/top")
    end_time = time.time()
    print(f"GetTopPlayers API call took: {end_time - start_time:.4f} seconds")
    return response.json()

# Fetch user rank
def get_user_rank(user_id):
    start_time = time.time()
    response = requests.get(f"{API_BASE_URL}/rank/{user_id}")
    end_time = time.time()
    print(f"GetPlayerRank API call took: {end_time - start_time:.4f} seconds")
    return response.json()

if __name__ == "__main__":
    while True:
        user_id = random.randint(1, 1000000)
        submit_score(user_id)
        print(get_top_players())
        print(get_user_rank(user_id))
        time.sleep(random.uniform(0.5, 2)) # Simulate real user interaction
