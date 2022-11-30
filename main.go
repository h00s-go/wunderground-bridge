package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request at %v\n: %v", time.Now(), r.URL.Path)
	url := fmt.Sprintf("http://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?%v", r.URL.RawQuery)
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weatherstation/updateweatherstation.php", handleUpdate)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
