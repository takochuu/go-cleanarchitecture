package interfaces

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/src/usecases"
)

type UserInterface interface {
	Create() (*usecases.User, error)
}

type UserHandler struct {
	UserInterface UserInterface
}

func (h UserHandler) Create(res http.ResponseWriter, req *http.Request) {
	h.UserInterface.Create()
}
