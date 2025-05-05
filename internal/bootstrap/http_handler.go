package bootstrap

import (
	"github.com/PongponZ/go-book-store/internal/delivery/http"
	"github.com/PongponZ/go-book-store/internal/delivery/http/handler"
)

func initializeHTTPHandler(u *Usecase) *http.Handlers {
	return &http.Handlers{
		UserHandler: handler.NewUserHandler(u.UserUsecase),
	}
}
