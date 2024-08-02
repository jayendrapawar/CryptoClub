// redux/cryptoSlice.ts
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

interface CryptoData {
  code: string;
  rate: number;
}

export const fetchCryptoData = createAsyncThunk<
  CryptoData[],  // The type of the data returned from the thunk
  void,          // The type of the argument passed to the thunk (none in this case)
  { rejectValue: string } // Optional: custom error handling type
>('crypto/fetchCryptoData', async (_, { rejectWithValue }) => {
  try {
    const response = await axios.get(process.env.NEXT_PUBLIC_API_URL as string); 
    return response.data;
  } catch (error) {
    return rejectWithValue('Failed to fetch crypto data');
  }
});

interface CryptoState {
  data: CryptoData[];
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null; // Add error state
}

const initialState: CryptoState = {
  data: [],
  status: 'idle',
  error: null,
};

const cryptoSlice = createSlice({
  name: 'crypto',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchCryptoData.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCryptoData.fulfilled, (state, action) => {
        state.status = 'succeeded';
        // Create a map for quick lookups
        const dataMap = new Map<string, number>();
        state.data.forEach(item => dataMap.set(item.code, item.rate));

        // Update the map with the new data
        action.payload.forEach(item => {
          dataMap.set(item.code, item.rate);
        });

        // Convert map back to array and sort by timestamp if needed
        state.data = Array.from(dataMap, ([code, rate]) => ({ code, rate }));
        state.error = null;
      })
      .addCase(fetchCryptoData.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload as string; // Handle errors
      });
  },
});

export default cryptoSlice.reducer;
