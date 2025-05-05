package response

import "github.com/PongponZ/go-book-store/internal/domain/entity"

type FindByID struct {
	Status string      `json:"status"`
	Data   entity.User `json:"data"`
}
