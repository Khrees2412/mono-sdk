package main

import (
	"encoding/json"
	"fmt"
)

type AccoutHolderReq struct {
	Entity     string `json:"entity"`
	First_name string
	Last_name  string
	Bvn        string
	Phone      string
	Address    struct {
		Address_line1 string
		Lga           string
		City          string
		State         string
	} `json:"address"`
}

func main() {

	acct := &AccoutHolderReq{
		Entity: "string",
		Address: struct {
			Address_line1 string
			Lga           string
			City          string
			State         string
		}{
			Address_line1: "adad",
		},
	}
	a, _ := json.Marshal(acct)
	fmt.Println(string(a))
}
