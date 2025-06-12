# Square Loyalty Backend

This is the **backend server** for the Square Loyalty Points project. It acts as a middle layer between the frontend application and Square’s API, handling:

- Customer creation and management  
- Loyalty account linking  
- Points earning and redemption  
- Loyalty event history retrieval  
- Input validation and API request formatting  

The backend is built using **Node.js** and **Express**, and communicates with Square’s APIs using their official SDK.

---

## Table of Contents

- [Features](#features)  
- [Tech Stack](#tech-stack)  
- [Setup & Installation](#setup--installation)  
- [Environment Variables](#environment-variables)  
- [Project Structure](#project-structure)  
- [Available Endpoints](#available-endpoints)  
- [Known Issues](#known-issues)  
- [Future Improvements](#future-improvements)  
- [Contact](#contact)  

---

## Features

- Create a customer using name and phone number  
- Create or fetch loyalty account for a customer  
- Earn and redeem loyalty points  
- Fetch current points balance  
- Fetch transaction/event history  
- Handle errors and invalid requests gracefully  

---

## Tech Stack

- **Node.js** – Runtime environment  
- **Express.js** – Web framework  
- **Square Node.js SDK** – Square API integration  
- **dotenv** – Environment variable management  
- **cors** – Enable Cross-Origin Resource Sharing  
- **nodemon** – Development server auto-restart  

---

## Setup & Installation

### Prerequisites

- Node.js (v16 or above)  
- A Square developer account and access to your **Sandbox credentials**

---

### Steps

1. Clone the repo:

   ```bash
   git clone <your-backend-repo-url>
   cd backend

2. Install dependencies:

    npm install

3. Create a .env file with your Square API credentials:

    PORT=8080
    SQUARE_ACCESS_TOKEN=YOUR_SANDBOX_ACCESS_TOKEN
    SQUARE_ENVIRONMENT=sandbox
    LOCATION_ID=YOUR_SANDBOX_LOCATION_ID

4. Start the development server:

    go run main.go