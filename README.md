## Overview
This repository is sample microservice exercise that contains multiple Go services that communicate with each other:
\- `user-service`
\- `catalogue-service`
\- `order-service`
\- `gateway`

## Prerequisites
\- Go 1.20+  

## Project Structure
\- `user-service`: Handles user-related operations.  
\- `catalogue-service`: Manages product catalog.  
\- `order-service`: Manages orders.  
\- `gateway`: GraphQL gateway for interacting with all services from one endpoint.

## Getting Started
1. Clone the repository.
2. Navigate to each service folder and run `go build .` to compile.

## Running Locally
1. Start each service on different ports using `go run .` in their respective folders.
2. Go to the `gateway` folder and run `go run .` to serve a single GraphQL endpoint.

## Contributing
1. Fork the repository.
2. Create a new branch.
3. Make changes and submit a pull request.

## License
Distributed under an open-source license. See `LICENSE` for details.