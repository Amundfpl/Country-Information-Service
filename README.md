# Assignment_1 - Country Information API

##  Overview
This project provides a **RESTful API** that retrieves country-related information, including **general details, population statistics, and API service status.** It is structured using best practices to maintain **scalability, readability, and modularity.**

---

## Project Structure
```
Assignment_1/
│── cmd/               #  Application entry points
│   ├── server/
│   │   ├── main.go    # Starts the server
│
│── internal/          #  Internal code (not importable by other projects)
│   ├── server/        # Server-related logic
│   │   ├── server.go  # Initializes the server
│   │   ├── router.go  # Registers all API routes
│
│── handlers/          #  Request handlers (business logic entry point)
│   ├── country_handler.go  # Handles country information requests
│   ├── population_handler.go  # Handles population-related requests
│   ├── status_handler.go  # Handles API status requests
│
│── models/            #  Data models (structs for JSON parsing)
│   ├── country.go  # Structs for country information
│   ├── population.go  # Structs for population responses
│   ├── status.go  # Structs for API status response
│
│── services/          #  Business logic (core functionality)
│   ├── country_service.go  # Fetches country info from APIs
│   ├── population_service.go  # Fetches population data from APIs
│   ├── status_service.go  # Retrieves API status and uptime
│
│── utils/             #  Utility functions (helpers & common functions)
│   ├── response.go  # Helper functions for handling responses
│
│── go.mod             #  Go module file
│── README.md          #  Documentation
```

---

##  Getting Started
###  Clone the Repository
```sh
 git clone https://github.com/your-username/Assignment_1.git
 cd Assignment_1
```

###  Install Dependencies
```sh
 go mod tidy
```

###  Set Up Environment Variables
Create a **.env** file in the root directory and add API URLs:
```sh
PORT=8080
REST_COUNTRIES_API=http://129.241.150.113:8080/v3.1
COUNTRIES_NOW_API=http://129.241.150.113:3500/api/v0.1
```

### Run the Server
```sh
 go run cmd/server/main.go
```

---

## API Endpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/countryinfo/v1/info/{countryCode}` | GET | Retrieve general country information |
| `/countryinfo/v1/population/{countryCode}` | GET | Get historical population data |
| `/countryinfo/v1/status` | GET | Check API status and uptime |

---

## Project Best Practices
✔ **Separation of Concerns:** Routes, handlers, services, and models are kept in separate directories.  
✔ **Scalability:** The modular structure allows easy expansion.  
✔ **Code Reusability:** Services handle business logic, making handlers simpler.  
✔ **Readability:** The project follows clean code principles for better maintainability.

---

## Contributing
Feel free to fork this repository and submit pull requests with improvements or bug fixes.

---

## License
This project is licensed under the MIT License - see the LICENSE file for details.

---

