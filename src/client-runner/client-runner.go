package main

import (
	EClient "github.com/ilivestrong/proto-buff-first-steps/src/client"
	"github.com/ilivestrong/proto-buff-first-steps/src/pb/mypbs"
)

func CreateEmployee() {
	EClient.CreateEmployee(&mypbs.EmployeeRequest{
		Name: "Deepak Pathak",
		Addresses: []*mypbs.Address{{
			Street:   "Ratchadaphisek Road",
			Postcode: 10600,
			Building: "Whizdom station ratchada thapra",
		}},
		Project: &mypbs.Project{
			Name:     "Telcomm",
			Position: "Senior Backend Engineer - Go",
			Skills: []*mypbs.Skill{
				{
					Name:       "Go",
					ExpInYears: 0,
				},
			},
		},
	})
}

func main() {
	CreateEmployee()
}
