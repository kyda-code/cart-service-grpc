package main

import (
	pb "cart-service/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	// defaultName is the default name used in the greeting.
	defaultName = "world"
)

var (
	// addr is the address of the server.
	addr = flag.String("addr", "localhost:50051", "the address of the server")
	//	name is the name to greet.
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := pb.NewCartServiceClient(conn)

	// Contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add an item to the cart
	r, err := c.AddItem(ctx, &pb.AddItemRequest{CartId: "1", Item: &pb.Item{Id: "1", Name: "apple", Quantity: 1, Price: 10.0}})
	if err != nil {
		log.Fatalf("could not add item: %v", err)
	}
	log.Printf("Cart: %v", r.Cart)

	// Get an item from the cart
	item, er := c.GetCart(ctx, &pb.GetCartRequest{CartId: "1"})
	if er != nil {
		log.Fatalf("could not get cart: %v", er)
	}
	log.Printf("Cart: %v", item.Cart)
}
