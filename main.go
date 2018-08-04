package main

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/src/interfaces"
)

func main() {

	h := interfaces.UserHandler{}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h.Create(res, req)
	})
}
