package service

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ilivestrong/proto-buff-first-steps/src/pb/mypbs"
	"google.golang.org/protobuf/proto"
)

type EmployeeService struct {
	mypbs.UnimplementedEmployeeServiceServer
}

func (s *EmployeeService) AddEmployee(ctx context.Context, e *mypbs.EmployeeRequest) (*mypbs.AddEmployeeResponse, error) {
	data, err := proto.Marshal(e)
	if err != nil {
		return nil, err
		// log.Fatal("Failed to serialize employee")
	}

	werr := ioutil.WriteFile(fmt.Sprintf("%s.txt", e.Name), data, 0644)
	if werr != nil {
		return nil, werr
		// log.Fatal("Failed to write employee to disk")
	}
	fmt.Println("Employee added successfully.")
	return &mypbs.AddEmployeeResponse{
		ID:      1,
		Success: true,
	}, nil
}
