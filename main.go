package main

import (
	"net/http"

	"github.com/fengbd/gopack/router"
)

func main() {
	http.ListenAndServe(":8080", router.Router)
}
