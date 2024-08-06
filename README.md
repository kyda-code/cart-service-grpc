# Cart Service

This project is a gRPC-based Cart Service implemented in Go. It allows clients to manage a shopping cart by adding, removing items, and retrieving the cart details.

## Prerequisites

- Go 1.22 or later
- Protocol Buffers compiler (`protoc`)
- Go plugins for gRPC

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/cart-service.git
   cd cart-service
   ```

2. **Install dependencies**:
   ```sh
   go mod tidy
   ```

3. **Generate gRPC code**:
   ```sh
   protoc --go_out=. --go-grpc_out=. proto/cart.proto
   ```

## Usage

### Running the Server

To start the gRPC server, run the following command:
```sh
go run server/main.go
```
The server will start and listen on port `50051`.

### Running the Client

To run the client and interact with the server, use the following command:
```sh
go run client/main.go
```

### Example Client Code

The client code demonstrates how to add an item to the cart, retrieve the cart, and remove an item from the cart.

```go
package main

import (
 "context"
 "log"
 "time"

 "google.golang.org/grpc"
 pb "cart-service/proto"
)

func main() {
 conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
 if err != nil {
  log.Fatalf("did not connect: %v", err)
 }
 defer conn.Close()
 client := pb.NewCartServiceClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 addItemResponse, err := client.AddItem(ctx, &pb.AddItemRequest{
  CartId: "cart1",
  Item: &pb.Item{
   Id:       "item1",
   Name:     "Apple",
   Quantity: 3,
   Price:    1.99,
  },
 })
 if err != nil {
  log.Fatalf("could not add item: %v", err)
 }
 log.Printf("Added Item: %v", addItemResponse.Cart)

 getCartResponse, err := client.GetCart(ctx, &pb.GetCartRequest{CartId: "cart1"})
 if err != nil {
  log.Fatalf("could not get cart: %v", err)
 }
 log.Printf("Cart: %v", getCartResponse.Cart)

 removeItemResponse, err := client.RemoveItem(ctx, &pb.RemoveItemRequest{
  CartId: "cart1",
  ItemId: "item1",
 })
 if err != nil {
  log.Fatalf("could not remove item: %v", err)
 }
 log.Printf("Removed Item: %v", removeItemResponse.Cart)
}
```

## Proto File

The `proto/cart.proto` file defines the gRPC service and messages used in the project.

```proto
syntax = "proto3";

package cart;
option go_package = "proto/cart";

// CartService is a service for managing a shopping cart.
service CartService {
  rpc AddItem(AddItemRequest) returns (AddItemResponse) {}
  rpc RemoveItem(RemoveItemRequest) returns (RemoveItemResponse) {}
  rpc GetCart(GetCartRequest) returns (GetCartResponse) {}
}

// Cart represents a shopping cart.
message Cart {
  string id = 1;
  repeated Item items = 2;
}

// Item represents an item in a shopping cart.
message Item {
  string id = 1;
  string name = 2;
  int32 quantity = 3;
  float price = 4;
}

// AddItemRequest represents a request to add an item to a cart.
message AddItemRequest {
  string cart_id = 1;
  Item item = 2;
}

// AddItemResponse represents a response to adding an item to a cart.
message AddItemResponse {
  Cart cart = 1;
}

// RemoveItemRequest represents a request to remove an item from a cart.
message RemoveItemRequest {
  string item_id = 1;
  string cart_id = 2;
}

// RemoveItemResponse represents a response to removing an item from a cart.
message RemoveItemResponse {
  Cart cart = 1;
}

// GetCartRequest represents a request to get a cart.
message GetCartRequest {
  string cart_id = 1;
}

// GetCartResponse represents a response to getting a cart.
message GetCartResponse {
  Cart cart = 1;
}
```

## Dependencies

The project uses the following dependencies:

- `github.com/golang/protobuf v1.5.4`
- `google.golang.org/grpc v1.65.0`
- `golang.org/x/net v0.26.0` (indirect)
- `golang.org/x/sys v0.21.0` (indirect)
- `golang.org/x/text v0.16.0` (indirect)
- `google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117` (indirect)
- `google.golang.org/protobuf v1.34.2` (indirect)

## License

This project is licensed under the MIT License.
```
