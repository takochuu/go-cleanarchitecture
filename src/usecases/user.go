package usecases

type UserUseCase struct{}

type User struct {
	ID int
}

func (u UserUseCase) Create() error {
	return nil
}
