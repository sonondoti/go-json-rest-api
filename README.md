go-json-rest-api

A REST API service built with Go 1.22 for managing cars, furniture, and flowers.
The project uses a JSON file for data storage and supports CRUD operations via HTTP methods (GET, POST, PUT, PATCH, DELETE). This project was developed as part of a test assignment for an IT internship.

Project Author
Alexander Dubin
Student at UlSU: PRI-О-22/1

Features

Entities:

        Cars: ID, brand name, model, mileage, and number of owners.
        Furniture: Name, manufacturer, height, width, length.
        Flowers: Flower name, quantity, price, delivery date.

CRUD Operations:

        POST: Create new records.
        GET: Retrieve lists or individual records by ID.
        PUT: Replace existing records.
        PATCH: Update partial data in records.
        DELETE: Remove records

Prerequisites

1.Install Go: Version 1.22 or later.
Download Go and follow the installation instructions for your operating system.

2.Clone the repository:

        git clone https://github.com/sonondoti/go-json-rest-api.git
        cd go-json-rest-api
        
3.Install dependencies:
Run go mod tidy to install necessary packages.

Setup Instructions

1.Run the server:
  Execute the following command in the project directory:
  go run main.go

2.Access the API: 
  By default, the server will run on http://localhost:9090. Use tools like Postman or curl to interact with the API.

API Endpoints
Cars:

  POST /cars – Add a new car.
  GET /cars – Retrieve all cars.
  GET /cars/{id} – Retrieve a specific car by ID.
  PUT /cars/{id} – Replace a car by ID.
  PATCH /cars/{id} – Update a car's fields.
  DELETE /cars/{id} – Remove a car by ID.

Furniture:

  POST /furniture – Add a new furniture item.
  GET /furniture – Retrieve all furniture items.
  GET /furniture/{id} – Retrieve a specific furniture item by ID.
  PUT /furniture/{id} – Replace a furniture item by ID.
  PATCH /furniture/{id} – Update a furniture item's fields.
  DELETE /furniture/{id} – Remove a furniture item by ID.


Flowers:
  POST /flowers – Add a new flower batch.
  GET /flowers – Retrieve all flower batches.
  GET /flowers/{id} – Retrieve a specific flower batch by ID.
  PUT /flowers/{id} – Replace a flower batch by ID.
  PATCH /flowers/{id} – Update a flower batch's fields.
  DELETE /flowers/{id} – Remove a flower batch by ID.

Testing

You can test the API using Postman, curl, or other HTTP clients.

Example curl requests:
  Add a new car:
    
    curl -X POST http://localhost:9090/cars 
    -H "Content-Type: application/json" 
    -d '{"brand":"Toyota", "model":"Camry", "mileage":15000, "owners":1}'

  Add a new furniture item:
    
    curl -X POST http://localhost:9090/furniture 
    -H "Content-Type: application/json" 
    -d '{"name":"Table", "manufacturer":"IKEA", "height":75, "width":120, "length":60}'

  Add a new flower batch:
    
    curl -X POST http://localhost:9090/flowers 
    -H "Content-Type: application/json" 
    -d '{"name":"Roses", "quantity":50, "price":200, "delivery_date":"2025-01-01"}'



