// src/components/BlogCard.tsx
import React from 'react';

interface BlogCardProps {
  currency: string;
  price: number;
}

const BlogCard: React.FC<BlogCardProps> = ({ currency, price }) => {
  return (
    <div className="max-w-sm rounded overflow-hidden shadow-lg m-4 bg-white">
      <div className="px-6 py-4">
        <div className="font-bold text-xl mb-2">{currency}</div>
        <p className="text-gray-700 text-base">
          Price: {price.toFixed(2)}
        </p>
      </div>
    </div>
  );
};

export default BlogCard;
