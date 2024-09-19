# Cards API

## Overview

The Cards API is a service designed to securely manage credit and debit cards within a vault. 
Its primary objective is to provide robust functionality for the creation, secure storage, and management of card information. 
This API ensures the protection of sensitive data through advanced security measures, offering a comprehensive and efficient interface for all card-related operations.

## Getting Started

### Prerequisites

- Docker
- Postman (for testing)

### Environment Variables

Create a `.env` file in the root directory with the following variables.

```plaintext
JWT_SECRET_KEY=ac3125206d48180fe7e3cb46f7f9f5252a0afbcc9bf05be66a7fa6d36958e8df
JWT_EXPIRATION_HOURS=24

POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=test
```

## Running the Application

To start the application, run the following command.

```
docker-compose up
```

## Endpoints

### Authentication

`
POST /login
`
Provides a JWT Token for authenticated requests.

### Cards

`
POST /cards-api/v1/cards
`
Adds a new card.

`
GET /cards-api/v1/cards/:id
`
Retrieves a card by its id.

`
PUT /cards-api/v1/cards/:id
`
Updates card information.

`
DELETE /cards-api/v1/cards/:id
`
Deletes a card by its id.

### Events

`
POST /cards-api/v1/events/update-cards
`
Updates multiple cards concurrently.

## Testing

You can import the Postman Collection for this API from the following link.

[Cards Postman Collection](https://www.postman.com/matteocarranza/workspace/cards/collection/8242347-5b358df7-1351-425b-901e-21085d766fb9?action=share&creator=8242347)
