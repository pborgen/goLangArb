package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/paulborgen/goLangArb/api/protos"
	interceptor "github.com/paulborgen/goLangArb/internal/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
    port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
    pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("Received: %v", in.GetName())
    return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
    flag.Parse()

    // Load TLS credentials
    creds, err := credentials.NewServerTLSFromFile("server-cert.pem", "server-key.pem")
    if err != nil {
        log.Fatalf("Failed to load TLS keys: %v", err)
    }

    // Create the auth interceptor
	jwtKey := []byte("your-secret-key") // Use a secure key in production
	authInterceptor := interceptor.NewAuthInterceptor(jwtKey)
    
    // Create gRPC server with TLS
    s := grpc.NewServer(
        grpc.Creds(creds),
        grpc.UnaryInterceptor(authInterceptor.Unary()),
    )

    pb.RegisterGreeterServer(s, &server{})


    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    //s := grpc.NewServer()
    //pb.RegisterGreeterServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
