package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"time"

	pb "github.com/paulborgen/goLangArb/api/protos"
	"github.com/paulborgen/goLangArb/internal/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
    defaultName = "world"
)

var (
    addr = flag.String("addr", "localhost:50051", "the address to connect to")
    name = flag.String("name", defaultName, "Name to greet")
)

func main() {
    flag.Parse()
	cert, err := tls.LoadX509KeyPair("../server/server-cert.pem", "../server/server-key.pem")
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}
	
	config := &tls.Config{
		InsecureSkipVerify: true,  // Disables hostname verification
		Certificates: []tls.Certificate{cert},
	}
	
    // Set up a connection to the server.
	creds := credentials.NewTLS(config)
    if err != nil {
        log.Fatalf("Failed to load credentials: %v", err)
    }

	 // Create connection with TLS
	conn, err := grpc.NewClient(*addr, 
        grpc.WithTransportCredentials(creds))
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()


    c := pb.NewGreeterClient(conn)

	// Generate a JWT token (in a real scenario, you'd get this from a login endpoint)
	token, err := auth.GenerateJWT("username")
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

    // Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

	// Add the token to the context metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
    
	if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    
	log.Printf("Greeting: %s", r.GetMessage())
}
