package main

import (
	"fmt"
	"net/http"

	r "github.com/ijasmoopan/usermanagement-api/router"
)
func main() {

	fmt.Println("API is starting..")

	http.ListenAndServe(":8080", r.Router())
}