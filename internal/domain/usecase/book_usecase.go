package usecase

import (
	"fmt"

	"github.com/PongponZ/go-book-store/internal/domain/constant"
)

type IBookUsecase interface{}

type bookUsecase struct{}

func NewBookUsecase() IBookUsecase {
	return &bookUsecase{}
}

func (u *bookUsecase) SomeFunc1(bookType string) {
	if bookType == constant.BookTypeFiction {
		fmt.Println("hi, SomeFunc1")
	}
}

func (u *bookUsecase) SomeFunc2(bookType string) {
	if bookType == constant.BookTypeFiction {
		fmt.Println("hi, SomeFunc2")
	}
}
