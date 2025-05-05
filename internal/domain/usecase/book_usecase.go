package usecase

import (
	"fmt"
)

type IBookUsecase interface{}

type bookUsecase struct{}

func NewBookUsecase() IBookUsecase {
	return &bookUsecase{}
}

func (u *bookUsecase) SomeFunc1(bookType string) {
	if bookType == "fiction" {
		fmt.Println("hi, SomeFunc1")
	}
}

func (u *bookUsecase) SomeFunc2(bookType string) {
	if bookType == "fiction" {
		fmt.Println("hi, SomeFunc2")
	}
}
