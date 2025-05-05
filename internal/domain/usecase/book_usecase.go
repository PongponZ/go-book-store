package usecase

type IBookUsecase interface{}

type bookUsecase struct{}

func NewBookUsecase() IBookUsecase {
	return &bookUsecase{}
}
