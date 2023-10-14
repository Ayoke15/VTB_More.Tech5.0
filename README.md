

Welcome to MyAPI documentation. This API is designed to fetch recommended offices and ATMs based on the user's location and preferences.

## 📌 Table of Contents
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
  - [Fetch Recommended Offices](#fetch-recommended-offices)
  - [Fetch Recommended ATMs](#fetch-recommended-atms)
- [Contributions and Feedback](#contributions-and-feedback)
- [License](#license)

## 🚀 Getting Started

To start using the API, ensure you have the necessary dependencies installed and run the application using:

```bash
python API.py
The API will be up and running at http://127.0.0.1:8080/.

📡 API Endpoints
Fetch Recommended Offices
Endpoint: /get_offices
Method: POST
Payload:

json
Copy code
{
    "user_location": [latitude, longitude],
    "user_filters": {...}
}
Response:
Returns the top 5 recommended offices based on the user's location and filters.

json
Copy code
[
    {"id": office_id1, "score": score1},
    ...
]
⚠️ Errors:

400 Bad Request: If no matching offices are found.
Fetch Recommended ATMs
Endpoint: /get_atms
Method: POST
Payload:

json
Copy code
{
    "user_location": [latitude, longitude],
    "user_filters": {...}
}
Response:
Returns the top 5 recommended ATMs based on the user's location and filters.

json
Copy code
[
    {"id": atm_id1, "score": score1},
    ...
]
⚠️ Errors:

400 Bad Request: If no matching ATMs are found.
💌 Contributions and Feedback
We welcome your feedback and contributions. Please feel free to open an issue or submit a pull request.

📜 License
This project is licensed under the MIT License. See the LICENSE.md file for details.