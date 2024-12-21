package main

import (
	"fmt"
	"log"
	"net/http"
)

var idx = 0

func homePage(w http.ResponseWriter, r *http.Request) {
	idx++
	fmt.Println(idx)
}

func apiRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":20086", nil))
}

func main() {
	apiRequests()
}
