# API Documentation

This document provides detailed information about the available API endpoints and their usage.

## Endpoints

### Products API

#### Get All Products
- **Method**: GET
- **Path**: /api/products
- **Description**: Retrieves a list of all products
- **Response**: List of Product objects wrapped in ApiResponse
```json
{
    "data": [
        {
            "id": int,
            "sku": string,
            "name": string,
            "description": string,
            "quantity": int,
            "unit_price": int
        }
    ],
    "message": string,
    "error": string
}
```

#### Get Product by ID
- **Method**: GET
- **Path**: /api/products/:id
- **Description**: Retrieves a specific product by ID
- **Parameters**: 
  - `id` (path parameter) - Product ID
- **Response**: Single Product object wrapped in ApiResponse
```json
{
    "data": {
        "id": int,
        "sku": string,
        "name": string,
        "description": string,
        "quantity": int,
        "unit_price": int
    },
    "message": string,
    "error": string
}
```

#### Create Product
- **Method**: POST
- **Path**: /api/products
- **Description**: Creates a new product
- **Request Body**:
```json
{
    "sku": string,
    "name": string,
    "description": string,
    "quantity": int,
    "unit_price": int
}
```
- **Response**: Created Product object wrapped in ApiResponse
```json
{
    "data": {
        "id": int,
        "sku": string,
        "name": string,
        "description": string,
        "quantity": int,
        "unit_price": int
    },
    "message": string,
    "error": string
}
```

#### Update Product
- **Method**: PUT
- **Path**: /api/products/:id
- **Description**: Updates an existing product
- **Parameters**: 
  - `id` (path parameter) - Product ID
- **Request Body**:
```json
{
    "sku": string,
    "name": string,
    "description": string,
    "quantity": int,
    "unit_price": int
}
```
- **Response**: Updated Product object wrapped in ApiResponse
```json
{
    "data": null,
    "message": string,
    "error": string
}
```

#### Delete Product
- **Method**: DELETE
- **Path**: /api/products/:id
- **Description**: Deletes a product
- **Parameters**: 
  - `id` (path parameter) - Product ID
- **Response**: Success message wrapped in ApiResponse
```json
{
    "data": null,
    "message": string,
    "error": string
}
```

## Response Format
All API endpoints return responses in the following format:
```json
{
    "data": any,
    "message": string,
    "error": string
}
```
