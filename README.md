# API Documentation

## Introduction

This API is designed to fetch recommended offices and ATMs based on the user's location and specified filters.

---

## Table of Contents

- [Introduction](#introduction)
- [Endpoints](#endpoints)
  - [Get Offices](#get-offices)
  - [Get ATMs](#get-atms)
- [Usage](#usage)
- [Responses](#responses)

---

## Endpoints

### Get Offices

**URL**: `/get_offices`

**Method**: `POST`

**Payload**:

```json
{
  "user_location": [latitude, longitude],
  "user_filters": {}
}
Description: Fetches the top 5 recommended offices based on the user's location and specified filters.

Get ATMs
URL: /get_atms

Method: POST

Payload:

json
Copy code
{
  "user_location": [latitude, longitude],
  "user_filters": {}
}
Description: Fetches the top 5 recommended ATMs based on the user's location and specified filters.

Usage
To use this API, make a POST request to the respective endpoint (/get_offices or /get_atms) with the required payload.

Responses
Responses will be returned in the following format:

For offices:

json
Copy code
[
  {
    "id": "office_id",
    "score": "score_value"
  },
  ...
]
For ATMs:

json
Copy code
[
  {
    "id": "atm_id",
    "score": "score_value"
  },
  ...
]
In case no offices or ATMs are found, a 400 Bad Request response will be returned with a corresponding message.

Note: Ensure to replace placeholders (like "latitude", "longitude", etc.) with actual data when making requests.

