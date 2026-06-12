package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write "Users List" to the ResponseWriter
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write "Products List" to the ResponseWriter
}

func main() {
	mux := http.NewServeMux()

	// TODO: Register usersHandler for the "/users" route
	// TODO: Register productsHandler for the "/products" route

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp1, err := http.Get(ts.URL + "/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp1.Body.Close()
	body1, _ := io.ReadAll(resp1.Body)
	fmt.Println(string(body1))

	resp2, err := http.Get(ts.URL + "/products")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(body2))
}
