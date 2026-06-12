package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	// TODO: Encode the user struct to JSON using json.Marshal
	// Print the JSON string using fmt.Println
	_ = user

	jsonStr := `{"id":2,"name":"Bob","email":"bob@example.com"}`
	var decoded User
	_ = jsonStr
	_ = decoded

	// TODO: Decode the JSON string into the decoded variable using json.Unmarshal
	// Print the name and email using:
	// fmt.Printf("Name: %s, Email: %s\n", decoded.Name, decoded.Email)
}

// Ensure imports are used (remove these once you complete the TODOs above)
var _ = json.Marshal
var _ = fmt.Println
