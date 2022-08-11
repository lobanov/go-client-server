package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lobanov/go-client-server/protocol"
)

type Options struct {
	ServerAddr string `short:"a" long:"address" description:"Server address to connect to" `
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)

		default:
			os.Exit(1)
		}
	}

	log.Printf("Connecting to %s", options.ServerAddr)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(options.ServerAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := pb.NewProductCatalogClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := client.FetchProducts(ctx, &pb.ProductsRequest{})
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}

	log.Println(res.Products)
}
