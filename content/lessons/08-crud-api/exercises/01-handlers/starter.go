package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write "Hello, API!" to the ResponseWriter using fmt.Fprint
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/hello")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(string(body))
}
