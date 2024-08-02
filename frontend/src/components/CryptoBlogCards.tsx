// src/components/CryptoBlogCards.tsx
import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { fetchCryptoData } from '../features/cryptoSlice';
import BlogCard from './BlogCard';

const CryptoBlogCards: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const top5Cryptos = useSelector((state: RootState) => {
    const data = state.crypto.data;
    // Sort by price in descending order and get top 5
    return [...data].sort((a, b) => b.rate - a.rate).slice(0, 5);
  });

  useEffect(() => {
    const intervalId = setInterval(() => {
      dispatch(fetchCryptoData());
    }, 5000); // Fetch data every 5 seconds

    return () => clearInterval(intervalId); // Cleanup interval on component unmount
  }, [dispatch]);

  return (
    <div className="flex flex-wrap justify-center">
      {top5Cryptos.map((crypto, index) => (
        <BlogCard key={index} currency={crypto.code} price={crypto.rate} />
      ))}
    </div>
  );
};

export default CryptoBlogCards;
