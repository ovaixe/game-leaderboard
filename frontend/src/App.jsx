import { useState, useEffect } from 'react';
import Leaderboard from './components/Leaderboard';
import UserRank from './components/UserRank';

function App() {
  const [leaderboardData, setLeaderboardData] = useState([]);
  const [userRankData, setUserRankData] = useState(null);
  const [userId, setUserId] = useState('');

  const fetchLeaderboard = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/leaderboard/top');
      const data = await response.json();
      setLeaderboardData(data.data);
    } catch (error) {
      console.error('Error fetching leaderboard:', error);
    }
  };

  const fetchUserRank = async () => {
    if (!userId) return;
    try {
      const response = await fetch(`http://localhost:8080/api/leaderboard/rank/${userId}`);
      const data = await response.json();
      setUserRankData(data.data);
    } catch (error) {
      console.error('Error fetching user rank:', error);
      setUserRankData(null);
    }
  };

  useEffect(() => {
    fetchLeaderboard();
    const interval = setInterval(fetchLeaderboard, 5000); // Poll every 5 seconds
    return () => clearInterval(interval);
  }, []);

  const handleUserIdChange = (e) => {
    setUserId(e.target.value);
  };

  const handleUserRankLookup = () => {
    fetchUserRank();
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex flex-col items-center py-10">
      <h1 className="text-5xl font-extrabold mb-10 text-blue-400">Game Leaderboard</h1>
      <div className="w-full max-w-4xl bg-gray-800 p-8 rounded-lg shadow-xl mb-10">
        <Leaderboard data={leaderboardData} />
      </div>
      <div className="w-full max-w-4xl bg-gray-800 p-8 rounded-lg shadow-xl">
        <h2 className="text-3xl font-bold mb-6 text-blue-300">User Rank Lookup</h2>
        <div className="flex items-center justify-center mb-6 space-x-4">
          <input
            type="text"
            placeholder="Enter User ID"
            value={userId}
            onChange={handleUserIdChange}
            className="p-3 rounded-md bg-gray-700 border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 flex-grow"
          />
          <button
            onClick={handleUserRankLookup}
            className="px-6 py-3 bg-blue-600 hover:bg-blue-700 rounded-md font-semibold transition duration-300 ease-in-out transform hover:scale-105"
          >
            Lookup Rank
          </button>
        </div>
        <UserRank data={userRankData} />
      </div>
    </div>
  );
}

export default App;
