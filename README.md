# Product API

This is a RESTful API for managing products, built using Go and the Gin framework, with SQLite as the database.

## Features

- Create, Read, Update, and Delete (CRUD) operations for products.
- Paginated product listing.
- Simple JSON-based API.

## Prerequisites

- Docker and Docker Compose installed on your machine.
- Basic knowledge of RESTful APIs and Go programming.

## Setup Instructions

- Clone Repo
- cd product-api
- docker-compose build
- docker-compose up

## Accessing the API
The API will be accessible at http://localhost:1721. Below are the available endpoints:

GET /products: Retrieve a list of products (supports pagination with limit and page query parameters).
GET /products/
: Retrieve a single product by ID.
POST /products: Create a new product. Requires a JSON body with name, description, and price.
PUT /products/
: Update an existing product by ID. Requires a JSON body with name, description, and price.
DELETE /products/
: Delete a product by ID.

## Example Requests
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "Product 1", "description": "Description of Product 1", "price": 19.99}'

curl -X GET http://localhost:1721/products 

curl -X DELETE http://localhost:1721/products \
-H "Content-Type: application/json" \
-d '{"name": "Product2", "description": "Description2", "price": 10.99}'

curl -X PUT http://localhost:1721/products \
-H "Content-Type: application/json" \
-d '{"name": "Product2", "description": "Description2", "price": 10.99}'