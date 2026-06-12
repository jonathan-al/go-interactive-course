package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  = make(map[int]User)
	nextID = 1
	mu     sync.Mutex
)

func createUser(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// 1. Decode the JSON request body into a User variable
	// 2. Lock the mutex
	// 3. Assign the user an ID from nextID, then increment nextID
	// 4. Store the user in the users map
	// 5. Unlock the mutex
	// 6. Set the Content-Type header to "application/json"
	// 7. Write status code 201 (http.StatusCreated)
	// 8. Encode the user as JSON to the response
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// 1. Extract the ID string by trimming "/users/" prefix from r.URL.Path
	// 2. Convert the ID string to int using strconv.Atoi
	//    If the conversion fails, write status 400 and return
	// 3. Lock the mutex and look up the user in the users map
	// 4. Unlock the mutex
	// 5. If the user is not found, write status 404 and return
	// 6. Set the Content-Type header to "application/json"
	// 7. Encode the user as JSON to the response
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", createUser)
	mux.HandleFunc("/users/", getUser)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	postBody := `{"name":"Alice","email":"alice@example.com"}`
	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(postBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("POST /users -> Status: %d\n", resp.StatusCode)
	fmt.Printf("Created: %s\n", strings.TrimSpace(string(body)))

	resp2, err := http.Get(ts.URL + "/users/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("GET /users/1 -> Status: %d\n", resp2.StatusCode)
	fmt.Printf("User: %s\n", strings.TrimSpace(string(body2)))
}

// Ensure imports are used (remove these blank identifiers once you complete the TODOs above)
var _ = strconv.Atoi
var _ = json.Marshal
