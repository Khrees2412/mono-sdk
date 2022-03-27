package utils

import "encoding/json"

// Convert string or interface{} to JSON
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
