// src/components/CryptoTable.tsx
import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

const CryptoTable: React.FC = () => {
  const cryptoData = useSelector((state: RootState) => state.crypto.data);
  const recentData = cryptoData.slice(-20); // Get the most recent 20 entries

  return (
    <div className="flex flex-col items-center mt-2">
      <div className="overflow-x-auto w-full">
      <h1 className="text-2xl font-bold mb-2 text-left">Watchlist</h1>
        <table className="min-w-full bg-white border border-gray-200 rounded-lg shadow-md">
          <thead>
            <tr className="bg-gray-100 border-b">
              <th className="py-2 px-4 text-left text-sm font-medium text-gray-600">Currency</th>
              <th className="py-2 px-4 text-left text-sm font-medium text-gray-600">Rate</th>
            </tr>
          </thead>
          <tbody>
            {recentData.map((data, index) => (
              <tr key={index} className="border-b">
                <td className="py-2 px-4 text-sm text-gray-800">{data.code}</td>
                <td className="py-2 px-4 text-sm text-gray-800">{data.rate.toFixed(2)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default CryptoTable;
