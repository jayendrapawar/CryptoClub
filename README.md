# CryptoClub

## Table of Contents

- [Technologies](#technologies)
- [Features](#features)
- [Setup](#setup)
    - [Backend Setup](#backend-setup)
    - [Frontend Setup](#frontend-setup)
- [Folder Structure](#folder-structure)

## Technologies

- **Frontend:**
    - Next.js
    - TypeScript
    - Redux Toolkit
    - Tailwind CSS
    - Axios

- **Backend:**
  - Go
  - MongoDB
  - Docker
  - go.mongodb.org/mongo-driver
  - github.com/joho/godotenv

## Features

- Fetches real-time cryptocurrency data from an external API.
- Stores the data in MongoDB.
- Displays the top 5 cryptocurrencies in a blog card format.
- Updates the frontend every 5 seconds to reflect the latest data.

## Setup

### Backend Setup

1. **Clone the repository:**
   ```sh
   https://github.com/jayendrapawar/CryptoClub.git
   cd CryptoClub

#### Navigate to the backend folder:

 - cd backend

#### Create a .env file and add your environment variables:

API_URL=your_api_url
  - API_URL=your_api_url
  - API_KEY=your_api_key
  - MONGO_URI=your_mongodb_uri
  - DB_NAME=your_database_name
  - COLLECTION_NAME=your_collection_name

#### Build and run the backend:

 - go mod tidy
 - go run src/main.go


### Frontend Setup

 - cd frontend

#### Create a .env.local file and add your environment variables:

 - NEXT_PUBLIC_API_URL=http://localhost:8080/api/recent

#### Install the dependencies:

 - npm install

#### Run the frontend:

 - npm run dev

## Folder Structure

```sh
CryptoClub/
│
├── .env
├── go.mod
├── go.sum
├── backend/
│   ├── src/
│   │   ├── main.go
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── fetch/
│   │   │   └── fetch.go
│   │   ├── store/
│   │   │   └── store.go
│   │   └── models/
│   │       └── models.go
│   ├── Dockerfile
│   └── .env
│
├── frontend/
│   ├── .env.local
│   ├── package.json
│   ├── tsconfig.json
│   ├── next.config.js
│   ├── public/
│   ├── pages/
│   │   ├── index.tsx
│   │   └── _app.tsx
│   ├── components/
│   │   ├── CryptoBlogCards.tsx
│   │   ├── BlogCard.tsx
│   ├── redux/
│   │   ├── store.ts
│   │   ├── cryptoSlice.ts
│   └── styles/
│       └── globals.css
│
├── README.md

