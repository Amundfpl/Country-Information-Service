package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler serves a structured HTML page with API documentation
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	baseURL := "http://" + r.Host // Dynamically get host

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
			.code-block { background: #272822; color: #fff; padding: 10px; border-radius: 5px; }
		</style>
	</head>
	<body>
		<h1>Country Information API</h1>
		<p><strong>API Version:</strong> v1</p>
		<p>This API provides country information, population statistics, and service status.</p>

		<h2>Available Endpoints</h2>
		<ul>
			<li><a href="` + baseURL + `/countryinfo/v1/info/NO">/countryinfo/v1/info/{countryCode}</a> - Retrieve general country information</li>
			<li><a href="` + baseURL + `/countryinfo/v1/population/NO">/countryinfo/v1/population/{countryCode}</a> - Get historical population data</li>
			<li><a href="` + baseURL + `/countryinfo/v1/status">/countryinfo/v1/status</a> - Check API status and uptime</li>
		</ul>

		<h2>Example API Call</h2>
		<pre>curl -X GET "` + baseURL + `/countryinfo/v1/info/NO"</pre>

		<h2>Sample JSON Response</h2>
		<pre class="code-block">
{
	"name": "Norway",
	"continents": ["Europe"],
	"population": 5391369,
	"languages": {"nno":"Norwegian Nynorsk", "nob":"Norwegian Bokmål"},
	"borders": ["FIN","SWE","RUS"],
	"flag": "https://flagcdn.com/w320/no.png",
	"capital": "Oslo",
	"cities": ["Bergen", "Oslo", "Trondheim"]
}
		</pre>

		<h2>Notes</h2>
		<p>Ensure that API endpoints are correctly formatted. If issues arise, check the <a href="` + baseURL + `/countryinfo/v1/status">status endpoint</a> to verify external API availability.</p>
	</body>
	</html>
	`

	if _, err := fmt.Fprint(w, html); err != nil {
		http.Error(w, "Failed to load API documentation", http.StatusInternalServerError)
	}
}
