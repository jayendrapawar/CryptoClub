// src/pages/index.tsx
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { fetchCryptoData } from '../features/cryptoSlice';
import CryptoTable from '../components/CryptoTable';
import CryptoBlogCards from '../components/CryptoBlogCards';

const IndexPage = () => {
  const dispatch = useDispatch<AppDispatch>();
  const cryptoStatus = useSelector((state: RootState) => state.crypto.status);
  const cryptoError = useSelector((state: RootState) => state.crypto.error);

  useEffect(() => {
    const intervalId = setInterval(() => {
      dispatch(fetchCryptoData());
    }, 5000); // Fetch data every 5 seconds

    return () => clearInterval(intervalId); // Cleanup interval on component unmount
  }, [dispatch]);

  return (
    <div className="container mx-auto px-4">
      <h1 className="text-2xl font-bold mt-8 text-center">CryptoClub</h1>
      {cryptoStatus === 'loading' && <p className="text-center">Loading...</p>}
      {cryptoStatus === 'failed' && <p className="text-center text-red-500">Error: {cryptoError}</p>}
      {cryptoStatus === 'succeeded' && (
        <>
          <CryptoBlogCards />
          <CryptoTable />
        </>
      )}
    </div>
  );
};

export default IndexPage;
