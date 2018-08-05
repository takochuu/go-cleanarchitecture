package interfaces

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/src/domains"
)

type UserUseCase interface {
	Create() (*domains.User, error)
}

type UserHandler struct {
	UserUseCase UserUseCase
}

func (h UserHandler) Create(res http.ResponseWriter, req *http.Request) {
	h.UserUseCase.Create()
}
