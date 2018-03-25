package main

import (
	"net/http"

	"github.com/IdanLoo/gopack/router"
)

func main() {
	http.ListenAndServe(":8080", router.Router)
}
