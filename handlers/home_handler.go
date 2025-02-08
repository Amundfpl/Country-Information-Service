package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler serves a structured HTML page with API documentation
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Country Information API</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; padding: 20px; background-color: #f8f9fa; color: #333; }
			h1 { font-size: 24px; border-bottom: 2px solid #007BFF; padding-bottom: 10px; }
			h2 { font-size: 20px; margin-top: 20px; }
			pre { background: #f4f4f4; padding: 10px; border-radius: 5px; border-left: 4px solid #007BFF; overflow-x: auto; }
			ul { list-style-type: none; padding: 0; }
			li { margin-bottom: 8px; }
			a { color: #007BFF; text-decoration: none; font-weight: bold; }
			a:hover { text-decoration: underline; }
		</style>
	</head>
	<body>
		<h1>Country Information API</h1>
		<p>This API provides comprehensive information about countries, including general details, population statistics, and external service status.</p>

		<h2>Available Endpoints</h2>
		<ul>
			<li><a href="/countryinfo/v1/info/NO">/countryinfo/v1/info/{countryCode}</a> - Retrieve general country information</li>
			<li><a href="/countryinfo/v1/population/NO">/countryinfo/v1/population/{countryCode}</a> - Get historical population data</li>
			<li><a href="/countryinfo/v1/status">/countryinfo/v1/status</a> - Check API status and uptime</li>
		</ul>

		<h2>How to Use</h2>
		<p>You can interact with the API using a web browser or tools such as Postman and cURL.</p>
		<pre>curl -X GET "http://localhost:8080/countryinfo/v1/info/NO"</pre>

		<h2>Notes</h2>
		<p>Ensure that the API endpoints are correctly formatted. If issues arise, check the status endpoint to verify external API availability.</p>
	</body>
	</html>
	`

	fmt.Fprint(w, html)
}
