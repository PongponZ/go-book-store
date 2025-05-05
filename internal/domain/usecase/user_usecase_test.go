package usecase_test

import (
	"context"
	"testing"

	"github.com/PongponZ/go-book-store/internal/domain/entity"
	"github.com/PongponZ/go-book-store/internal/domain/repository/mocks"
	"github.com/PongponZ/go-book-store/internal/domain/usecase"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_FindByID(t *testing.T) {
	var (
		mockUserRepository *mocks.IUserRepository
		userUsecase        usecase.IUserUsecase

		findByIDResponse *entity.User
		findByIDError    error
	)

	beforeEach := func() {
		mockUserRepository = &mocks.IUserRepository{}
		userUsecase = usecase.NewUserUsecase(mockUserRepository)

		findByIDResponse = nil
		findByIDError = nil

		mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(
			func(context.Context, string) *entity.User {
				return findByIDResponse
			},
			func(context.Context, string) error {
				return findByIDError
			},
		)
	}

	t.Run("should call FindByID of user repository", func(t *testing.T) {
		beforeEach()

		userUsecase.FindByID(context.Background(), "123")

		mockUserRepository.AssertCalled(t, "FindByID", context.Background(), "123")
	})
}
