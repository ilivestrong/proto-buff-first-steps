package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ilivestrong/proto-buff-first-steps/src/pb/mypbs"
	"google.golang.org/protobuf/proto"
)

func main() {
	ep := mypbs.EmployeeRequest{
		Name:      "Deepak Pathak",
		Addresses: []*mypbs.Address{{Street: "Ratchadaphisek Road", Postcode: 10060, Building: "Whizdom Station Ratchada Thapra"}},
		Project: &mypbs.Project{
			Name:     "Telcomm",
			Position: "Senior Backend Engineer",
			Skills: []*mypbs.Skill{
				{
					Name:       "Go",
					ExpInYears: 0,
				},
			},
		},
	}

	data, err := proto.Marshal(&ep)
	if err != nil {
		log.Fatal("Failed to serialize employee")
	}

	werr := ioutil.WriteFile("employee1.txt", data, 0644)
	if werr != nil {
		log.Fatal("Failed to write employee to disk")
	}
	fmt.Println("Wrote employee message to a file")

	fmt.Println("Reading message from file...")
	in, rerr := ioutil.ReadFile("employee1.txt")
	if rerr != nil {
		log.Fatal("Failed to read message from file\n", rerr)
	}
	newEmp := &mypbs.EmployeeRequest{}
	proto.Unmarshal(in, newEmp)
	fmt.Printf("\nEmployee name: %s\nBuilding Name: %s\n", (*newEmp).Name, newEmp.GetAddresses()[0].GetBuilding())
	fmt.Printf("Project: %s, Skills: %s\n", newEmp.Project.Name, newEmp.Project.Skills[0].Name)
}
