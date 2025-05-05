package bootstrap

import "github.com/PongponZ/go-book-store/internal/domain/usecase"

type Usecase struct {
	UserUsecase usecase.IUserUsecase
}

func initializeUsecase(repository *Repository) *Usecase {
	userUsecase := usecase.NewUserUsecase(repository.UserRepository)

	return &Usecase{UserUsecase: userUsecase}
}
