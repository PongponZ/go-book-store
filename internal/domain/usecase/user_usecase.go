package usecase

import (
	"context"

	"github.com/PongponZ/go-book-store/internal/domain/entity"
	"github.com/PongponZ/go-book-store/internal/domain/repository"
)

type IUserUsecase interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	DeleteByID(ctx context.Context, id string) error
}

type userUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) IUserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) Create(ctx context.Context, user *entity.User) error {
	return u.userRepository.Create(ctx, user)
}

func (u *userUsecase) FindByID(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepository.FindByID(ctx, id)
}

func (u *userUsecase) Update(ctx context.Context, user *entity.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u *userUsecase) DeleteByID(ctx context.Context, id string) error {
	return u.userRepository.DeleteByID(ctx, id)
}
