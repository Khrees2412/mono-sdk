package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Convert string or interface{} to JSON
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetMonoId(code, user_id string) (interface{}, interface{}) {

	url := "https://api.withmono.com/account/auth"

	postBody, _ := json.Marshal(map[string]string{
		"code": code,
	})
	responseBody := []byte(postBody)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(responseBody))

	// req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err.Error(), nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	mono_id := PrettyPrint(body)

	return nil, mono_id

}
func main() {
	fmt.Println("App started")
}
