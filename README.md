# gopush

This repository contains a simple Go API with four endpoints for user authentication and user information retrieval.

## Setup
``` 
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/pusher.proto
```
```
local-env/docker-compose up -d
```

## Endpoints

1. **Landing Page**
    - Route: `/`
    - Description: Returns a simple string as a response.
    - HTTP Method: GET

2. **User Registration**
    - Route: `/api/auth/register`
    - Description: Registers a user by accepting an email and password in the request body as JSON.
    - HTTP Method: POST
    - Request Body:
      ```json
      {
        "email": "user@example.com",
        "password": "password123"
      }
      ```

3. **User Login**
    - Route: `/api/auth/login`
    - Description: Logs in a user by accepting an email and password in the request body as JSON. It returns a JSON token that can be used for authentication.
    - HTTP Method: POST
    - Request Body:
      ```json
      {
        "email": "user@example.com",
        "password": "password123"
      }
      ```
    - Response:
      ```json
      {
       "user": {
         "id": 1,
         "email": "user@example.com"
        },
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzX2F0IjoxNjg3MjQ5ODM2LCJ1c2VyIjp7ImlkIjoxLCJlbWFpbCI6ImFzZEBnbWFpbC5jb20ifX0.MRJETBn_ZprEbzMK4558C4ZO9J2RKgrOcLCqwrGnb1M",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzX2F0IjoxNjg4MzczMDM2LCJ1c2VyIjp7ImlkIjoxLCJlbWFpbCI6ImFzZEBnbWFpbC5jb20ifX0.9nVyOnER1ixLTxfN2KnAjqRg9DvzJ5jm10kT1Fu3iqk"
      }
      ```

4. **User Information**
    - Route: `/api/user`
    - Description: Returns the user information.
    - HTTP Method: GET
    - Authentication: Requires the "Authorization" header with the JSON token received from the login endpoint.
    - Example Request:
      ```
      GET /api/user
      Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
      ```
    - Example Response:
      ```json
      {
        "email": "user@example.com",
        "name": "John Doe",
      }
      ```
