package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%v]: %v\n", time.Now().Local(), r.URL.RawQuery)
	url := fmt.Sprintf("http://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?%v", r.URL.RawQuery)
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(url))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weatherstation/updateweatherstation.php", handleUpdate)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
