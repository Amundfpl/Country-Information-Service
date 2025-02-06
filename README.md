# Country-Information-Service

##  Project Overview
This project is a RESTful API service that provides country-related information using external APIs. The service fetches country details, population statistics, and API health status.

##  Supported Endpoints
###  General Country Information
- **GET** `/countryinfo/v1/info/{countryCode}`
    - Retrieves details such as country name, population, capital, languages, borders, flag, and a list of cities.
    - Supports an optional **`limit`** parameter to control the number of cities returned.

###  Population Data
- **GET** `/countryinfo/v1/population/{countryCode}`
    - Returns full historical population data.
- **GET** `/countryinfo/v1/population/{countryCode}?limit={startYear-endYear}`
    - Filters population history by a given time frame and computes the mean population.

###  API Health Check
- **GET** `/countryinfo/v1/status`
    - Provides the availability status of external APIs and the uptime of the service.

---

##  Request/Response Specification
For the specifications, the following syntax applies:

- **`{:value}`** indicates **mandatory** input parameters specified by the user.
- **`{value}`** indicates **optional** input specified by the user, where `value` can itself contain further optional input.
- The same notation applies for HTTP parameters:
    - **`{?:param}`** is a **mandatory** parameter.
    - **`{?param}`** is an **optional** parameter.

---

##  Setup and Usage
### **1️⃣ Prerequisites**
- Go **1.19+** installed.
- Internet connection to access external APIs.

### **2️⃣ Running the API Locally**
```sh
# Clone the repository
git clone https://github.com/amundfpl/country-info-service.git
cd country-info-service

# Run the application
go run main.go
```

### **3️⃣ Example API Requests**
```sh
# Fetch country information (default 10 cities)
curl -X GET "http://localhost:8080/countryinfo/v1/info/NO"

# Fetch country information with 5 cities
curl -X GET "http://localhost:8080/countryinfo/v1/info/NO?limit=5"

# Get population history from 2010 to 2015
curl -X GET "http://localhost:8080/countryinfo/v1/population/NO?limit=2010-2015"

# Check API status
curl -X GET "http://localhost:8080/countryinfo/v1/status"
```

---

##  Notes
- Ensure that the API endpoints you are calling match the **self-hosted versions** provided in the project documentation.
- Use **Postman** or another REST client to explore the API responses in more detail.

