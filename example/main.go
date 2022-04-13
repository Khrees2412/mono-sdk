package main

import (
	"fmt"

	mono "github.com/khrees2412/mono-sdk/lib"
)

func main() {
	mono, err := mono.NewClient("")
	if err != nil {
		fmt.Println(err)
	}
	mono.Connect.Unlink("")

}
