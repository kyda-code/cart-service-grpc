package main

import (
	pb "cart-service/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement the CartServiceServer interface.
type server struct {
	pb.UnimplementedCartServiceServer
	carts map[string]*pb.Cart
}

// newServer initializes a new server with an empty map to store carts.
func newServer() *server {
	return &server{
		carts: make(map[string]*pb.Cart),
	}
}

// AddItem adds an item to the cart. If the cart does not exist, it creates a new one.
func (s *server) AddItem(ctx context.Context, in *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	cart, exists := s.carts[in.CartId]
	if !exists {
		cart = &pb.Cart{Id: in.CartId}
		cart.Items = append(cart.Items, in.Item)
		s.carts[in.CartId] = cart
	}
	return &pb.AddItemResponse{Cart: cart}, nil
}

// GetCart retrieves the cart by its ID. If the cart does not exist, it returns nil.
func (s *server) GetCart(ctx context.Context, in *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	cart, exists := s.carts[in.CartId]
	if !exists {
		return &pb.GetCartResponse{Cart: nil}, nil
	}
	return &pb.GetCartResponse{Cart: cart}, nil
}

// RemoveItem removes an item from the cart by its ID. If the cart or item does not exist, it returns nil.
func (s *server) RemoveItem(ctx context.Context, in *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {
	cart, exists := s.carts[in.CartId]
	if !exists {
		return &pb.RemoveItemResponse{Cart: nil}, nil
	}
	newItems := make([]*pb.Item, 0)
	for _, item := range cart.Items {
		if item.Id != in.ItemId {
			newItems = append(newItems, item)
		}
	}
	cart.Items = newItems
	return &pb.RemoveItemResponse{Cart: cart}, nil
}

/*
// ClearCart clears all items from the cart. If the cart does not exist, it returns nil.
func (s *server) ClearCart(ctx context.Context, in *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
cart, exists := s.carts[in.CartId]
if !exists {
	return &pb.ClearCartResponse{Cart: nil}, nil
}
cart.Items = make([]*pb.Item, 0)
return &pb.ClearCartResponse{Cart: cart}, nil
}
*/

// main starts the gRPC server and listens on port 50051.
func main() {
	// Create a new gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCartServiceServer(s, newServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
