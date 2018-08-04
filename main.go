package main

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/src/interfaces"
	"github.com/takochuu/go-cleanarchitecture/src/usecases"
)

func main() {

	// TODO Router化
	h := interfaces.UserHandler{}
	h.UserInterface = &usecases.UserUseCase{}
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h.Create(res, req)
	})
}
