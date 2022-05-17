package main

import (
	"fmt"

	mono "github.com/khrees2412/mono-sdk/lib"
)

func main() {
	mono, err := mono.NewClient("test_sk_l5J0JlLVbRZMsCZzjD6b")
	if err != nil {
		fmt.Println(err)
	}
	id, e := mono.Connect.GetAccountId("")

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(id)

}
