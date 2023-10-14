# API Documentation for VTB_More.Tech5.0

## Table of Contents

1. [Introduction](#introduction)
2. [Endpoints](#endpoints)
   - [Get Offices](#get-offices)
   - [Get ATMs](#get-atms)
3. [Usage](#usage)
4. [Responses](#responses)

---

## Introduction

This API is specifically tailored to provide recommendations for offices and ATMs based on user input, which includes their current location and any specific filters they'd like to apply.

---

## Endpoints

### Get Offices

- **Endpoint**: `/get_offices`
- **HTTP Method**: POST
- **Payload**:
  ```json
  {
      "user_location": [latitude, longitude],
      "user_filters": {}
  }

Purpose: This endpoint will return the top 5 recommended offices based on the user's current location and filters.
Get ATMs
Endpoint: /get_atms
HTTP Method: POST
Payload:
{
    "user_location": [latitude, longitude],
    "user_filters": {}
}

Purpose: This endpoint will return the top 5 recommended ATMs based on the user's current location and filters.
Usage
To interact with the API:

Set up a POST request to the desired endpoint, either /get_offices for offices or /get_atms for ATMs.
In the body of your request, include a JSON payload with your current user_location and any user_filters you wish to apply.
The API will respond with the top 5 recommendations based on your criteria.
Responses
The API will return responses in the following formats:

For Offices:
[
    {
        "id": "office_id",
        "score": "score_value"
    },
    ...
]
For ATMs:
[
    {
        "id": "atm_id",
        "score": "score_value"
    },
    ...
]

Note: Replace placeholders (like "latitude", "longitude", etc.) with actual data when making requests.