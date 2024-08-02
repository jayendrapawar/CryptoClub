// src/features/rootReducer.ts
import { combineReducers } from '@reduxjs/toolkit';
import cryptoReducer from './cryptoSlice';

export const rootReducer = combineReducers({
  crypto: cryptoReducer,
});

export type RootState = ReturnType<typeof rootReducer>;
