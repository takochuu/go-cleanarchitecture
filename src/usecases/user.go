package usecases

import "github.com/takochuu/go-cleanarchitecture/src/domains"

type UserUseCase struct {
	repo UserInterface
}

type UserInterface interface {
	Create() (*domains.User, error)
}

func NewUserUseCase(r UserInterface) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) Create() (*domains.User, error) {
	user, err := u.repo.Create()
	if err != nil {
		return nil, err
	}
	return user, nil
}
