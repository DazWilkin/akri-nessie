package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/DazWilkin/akri-nessie/protos"

	"google.golang.org/grpc"
)

var (
	grpcEndpoint = flag.String("grpc_endpoint", "", "The endpoint of the gRPC server.")
)

func main() {
	log.Println("[main] Starting gRPC client")
	defer func() {
		log.Println("[main] Stopping gRPC client")
	}()

	flag.Parse()
	if *grpcEndpoint == "" {
		log.Fatal("[main] Unable to start client. Requires endpoint to a gRPC Server.")
	}

	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	log.Printf("Connecting to gRPC server [%s]", *grpcEndpoint)
	conn, err := grpc.Dial(*grpcEndpoint, dialOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewNessieClient(conn)
	ctx := context.Background()

	for {
		log.Println("[main:loop]")

		// Call GetNessieNow
		{
			rqst := &pb.NotifyRequest{}
			resp, err := client.GetNessieNow(ctx, rqst)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("[main:loop] Success: %+v", resp.GetFrame())
		}

		// Add a pause between iterations
		log.Println("[main:loop] Sleeping 5 seconds")
		time.Sleep(5 * time.Second)
		log.Println("[main:loop] Resuming")
	}

}
