package usecases

type UserUseCase struct{}

type User struct {
	ID int
}

func (u UserUseCase) Create() (*User, error) {
	return nil, nil
}
