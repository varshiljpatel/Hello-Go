package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const PORT int = 8000

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", helloHandler)

	go func() {
		err := http.ListenAndServe(":" + strconv.Itoa(PORT), nil)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server is listening on port ", PORT)
	select {} 
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var response Response = Response{
		Status:  200,
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}