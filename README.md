# test-dagangan-go-api

### Introduction
This API was built using GO and MongoDB, which provides CRUD operations, filtering, and pagination for a product database. This API project is designed to make it easy for developers to interact with the product database by providing a simple yet powerful interface.
### Features
* Basic CRUD functionality.
* Filtering data by categories.
* Data Pagination by manipulating offset and limit.
### Installation Guide
* Clone this repository [here](https://github.com/EdgarAllanPoo/test-dagangan-go-api.git).
* Run go mod download on your terminal from the root directory
* You should be working with your MongoDB database.
* Create an .env file in your project root folder and add your variables. See .env.sample for assistance.
### Usage
* Move into the src folder on your terminal.
* Run go run main.go to start the application.
* Connect to the API using Postman on port specified in .env.
### API Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| POST | /product | To create a new product instance |
| GET | /products | To retrieve all instances of product |
| GET | /products?category={category} | To retrieve all instances of product based on category |
| GET | /products?offset={offset}&limit={limit} | To retrieve all instances of product and paginate based on offset and limit |
| GET | /product/:Id | To retrieve one instance of product based on id |
| PUT | /product/:Id | To edit the details of one instance of product based on id |
| DELETE | /product/:Id | To delete one instance of product based on id |
### Technologies Used
* [GO](https://go.dev/)
* [MongoDB](https://www.mongodb.com/)
### Authors
* [EdgarAllanPoo](https://github.com/EdgarAllanPoo)
