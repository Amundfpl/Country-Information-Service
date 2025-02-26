# Assignment_1 - Country Information API

## Overview
This project provides a **RESTful API** to retrieve country-related information, including **general details, population statistics, and API service status**. The project follows best practices for **scalability, readability, and modularity**.

## Project Structure
```
Assignment_1/
├── .github/workflows/      
│   ├── sync.yaml           
│   ├── workflow.yaml
├── cmd/                    
│   ├── server/
│   │   ├── main.go
├── internal/               
│   ├── server/
│   │   ├── server.go       
│   │   ├── router.go       
│   ├── utils/
│   │   ├── constants.go    
│   │   ├── response.go
├── handlers/               
│   ├── country_handler.go   
│   ├── home_handler.go      
│   ├── population_handler.go  
│   ├── status_handler.go
├── models/                 
│   ├── country.go          
│   ├── population.go       
│   ├── status.go
├── pkg/                    
│   ├── services/
│   │   ├── country.go       
│   │   ├── population.go    
│   │   ├── status.go
├── .env                    
├── go.mod                  
├── README.md               
├── .gitignore              
```
## Getting Started

### 1. Clone the Repository
git clone https://github.com/YOUR_USERNAME/Assignment_1.git
cd Assignment_1

### 2. Install Dependencies
go mod tidy

### 3. Set Up Environment Variables
Create a **`.env`** file in the root directory and add:
PORT=8080
REST_COUNTRIES_API=http://129.241.150.113:8080/v3.1
COUNTRIES_NOW_API=http://129.241.150.113:3500/api/v0.1
DEFAULT_LIMIT=10

**Make sure to add `.env` to `.gitignore`** so it’s not committed to GitHub.

### 4. Run the Server
go run cmd/server/main.go

If `PORT` is **not set**, it will default to **8080**.

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/countryinfo/v1/info/{countryCode}` | GET | Retrieve general country information |
| `/countryinfo/v1/population/{countryCode}` | GET | Get historical population data |
| `/countryinfo/v1/status` | GET | Check API status and uptime |

## Best Practices
✔ **Uses `.env` for dynamic configuration** (port, API URLs, and limits).  
✔ **Separation of Concerns** → Routes, handlers, services, and models are structured separately.  
✔ **Scalability** → Modular structure allows easy expansion.  
✔ **Code Reusability** → Services handle business logic, making handlers simpler.  
✔ **Readability** → The project follows **clean code principles** for better maintainability.

## Contributing
Feel free to **fork** this repository and submit **pull requests** with improvements or bug fixes.


