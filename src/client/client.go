package EClient

import (
	"context"
	"fmt"
	"log"

	"github.com/ilivestrong/proto-buff-first-steps/src/pb/mypbs"
	"google.golang.org/grpc"
)

func CreateEmployee(e *mypbs.EmployeeRequest) {
	// var opts []grpc.DialOption
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect server: %v", err)
	}

	client := mypbs.NewEmployeeServiceClient(conn)

	if resp, err := client.AddEmployee(context.Background(), e); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("New Employee created, ID: %d", resp.ID)
	}
}
