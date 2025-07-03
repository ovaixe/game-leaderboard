import React from 'react';

const Leaderboard = ({ data }) => {
  return (
    <div className="leaderboard-container">
      <h2 className="text-3xl font-bold mb-6 text-blue-300">Top 10 Leaderboard</h2>
      {data.length === 0 ? (
        <p className="text-gray-400">Loading leaderboard...</p>
      ) : (
        <table className="min-w-full bg-gray-700 rounded-lg overflow-hidden shadow-lg">
          <thead className="bg-gray-600">
            <tr>
              <th className="py-3 px-4 text-left text-sm font-semibold text-gray-200 uppercase tracking-wider">Rank</th>
              <th className="py-3 px-4 text-left text-sm font-semibold text-gray-200 uppercase tracking-wider">Username</th>
              <th className="py-3 px-4 text-left text-sm font-semibold text-gray-200 uppercase tracking-wider">Score</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-600">
            {data.map((entry, index) => (
              <tr key={index} className="hover:bg-gray-600 transition duration-150 ease-in-out">
                <td className="py-3 px-4 whitespace-nowrap text-gray-300">{entry.rank}</td>
                <td className="py-3 px-4 whitespace-nowrap text-gray-300">{entry.username}</td>
                <td className="py-3 px-4 whitespace-nowrap text-gray-300">{entry.total_score}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default Leaderboard;