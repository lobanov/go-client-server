package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"

	pb "github.com/lobanov/go-client-server/protocol"
)

type Options struct {
	Port uint `short:"p" long:"port" description:"Port to listen to" default:"1234"`
}

type productCatalogServer struct {
	pb.UnimplementedProductCatalogServer

	products []*pb.Product
}

func (s *productCatalogServer) FetchProducts(context.Context, *pb.ProductsRequest) (*pb.ProductsResponse, error) {
	return &pb.ProductsResponse{Products: s.products}, nil
}

func newServer() *productCatalogServer {
	s := &productCatalogServer{}
	s.products = make([]*pb.Product, 0, 10)
	s.products = append(s.products, &pb.Product{Sku: "P001", Name: "Widget 1"})
	return s
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

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", options.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductCatalogServer(grpcServer, newServer())
	log.Printf("Listening on port %d", options.Port)
	grpcServer.Serve(lis)
}
