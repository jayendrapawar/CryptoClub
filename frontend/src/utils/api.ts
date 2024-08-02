// src/utils/api.ts
import axios from 'axios';

export const fetchRecentData = async () => {
  const response = await axios.get(process.env.NEXT_PUBLIC_API_URL as string);
  return response.data;
};
