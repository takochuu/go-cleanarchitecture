package main

import (
	"net/http"

	"github.com/takochuu/go-cleanarchitecture/src/infrastructure/repository"
	"github.com/takochuu/go-cleanarchitecture/src/interfaces"
	"github.com/takochuu/go-cleanarchitecture/src/usecases"
)

func main() {

	// TODO Router化
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		h := interfaces.UserHandler{}
		u := repository.NewUserRepository()
		h.UserUseCase = usecases.NewUserUseCase(u)
		h.Create(res, req)
	})
}
