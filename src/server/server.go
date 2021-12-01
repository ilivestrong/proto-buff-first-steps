package EServer

import (
	"fmt"
	"log"
	"net"

	"github.com/ilivestrong/proto-buff-first-steps/src/pb/mypbs"
	es "github.com/ilivestrong/proto-buff-first-steps/src/service"
	"google.golang.org/grpc"
)

const PORT = 50051

func BootServer() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatalf("Failed to boot server, error: %s", err.Error())
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	mypbs.RegisterEmployeeServiceServer(grpcServer, new(es.EmployeeService))
	fmt.Println("gRPC server running!")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %s", err.Error())
	}

}
