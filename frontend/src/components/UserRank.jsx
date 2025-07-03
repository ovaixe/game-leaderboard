import React from 'react';

const UserRank = ({ data }) => {
  if (!data) {
    return <p className="text-gray-400 text-center">Enter a User ID and click 'Lookup Rank' to see their rank.</p>;
  }

  return (
    <div className="bg-gray-700 p-6 rounded-lg shadow-md mt-6">
      <h3 className="text-2xl font-bold mb-4 text-blue-300">User Rank</h3>
      <p className="text-gray-300 mb-2"><strong className="text-blue-200">Username:</strong> {data.username}</p>
      <p className="text-gray-300 mb-2"><strong className="text-blue-200">Rank:</strong> {data.rank}</p>
      <p className="text-gray-300"><strong className="text-blue-200">Score:</strong> {data.total_score}</p>
    </div>
  );
};

export default UserRank;