# JIT Access Service

A simple gRPC service with HTTP gateway for handling Just-In-Time (JIT) access requests.

## Overview

This project implements an `AccessService` that allows users to request temporary access with specific roles and durations. The service runs both a gRPC server (port 50051) and an HTTP gateway (port 5000) using gRPC-Gateway.

## Output

<img width="1146" height="112" alt="image" src="https://github.com/user-attachments/assets/1f84d386-2bb9-410d-842f-fa727ccf446f" />


## Prerequisites

- Go 1.25.0 or later
- Protocol Buffers compiler (protoc)
- gRPC-Gateway plugins

## Installation

1. Clone the repository and navigate to the project directory.
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Generate protobuf files (if needed):
   ```bash
   protoc -I ./proto -I ./third_party --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./proto/jit.proto
   ```

## Running the Service

Start the server:
```bash
go run server/main.go
```

The service will start:
- gRPC server on `localhost:50051`
- HTTP gateway on `localhost:5000`

## API Endpoints

### Request Access
- **Method**: POST
- **URL**: `http://localhost:5000/v1/access/request`
- **Body**:
  ```json
  {
    "userId": "user123",
    "role": "admin",
    "durationMinutes": 60,
    "justification": "Need access for maintenance"
  }
  ```
- **Response**:
  ```json
  {
    "requestId": "req-123",
    "status": "approved"
  }
  ```

## Testing

Use the provided Postman collection (`jit-grpc-gateway.postman_collection.json`) to test the API endpoints.

## Project Structure

- `proto/jit.proto`: Protocol buffer definitions
- `server/main.go`: Main server implementation
- `third_party/`: Google API annotations
- Generated files: `jit.pb.go`, `jit_grpc.pb.go`, `jit.pb.gw.go`
