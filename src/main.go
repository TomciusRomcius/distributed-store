package main

import (
	"net/http"

	"github.com/tomciusromcius/distributed-store/src/controllers"
)

func main() {
	http.HandleFunc("/query", controllers.QueryController)
	http.ListenAndServe(":8080", nil)
}
