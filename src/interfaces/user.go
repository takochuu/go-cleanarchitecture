package interfaces

import "net/http"

type UserInterface interface {
	Create() error
}

type UserHandler struct {
	UserInterface UserInterface
}

func (h UserHandler) Create(res http.ResponseWriter, req *http.Request) {
}
