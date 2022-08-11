package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	pb "github.com/lobanov/go-client-server/protocol"
)

type Options struct {
	Port uint `short:"p" long:"port" description:"Port to listen to" default:"1234"`
}

type productCatalogServer struct {
}

func (s *productCatalogServer) FetchProducts(context.Context, *pb.ProductsRequest) (*pb.ProductsResponse, error) {

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
	fmt.Println(options)
}
