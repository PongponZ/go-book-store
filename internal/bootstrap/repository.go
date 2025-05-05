package bootstrap

import (
	"github.com/PongponZ/go-book-store/config"
	"github.com/PongponZ/go-book-store/internal/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UserRepository repository.IUserRepository
}

func initializeRepository(config *config.Config, mongoClient *mongo.Client) *Repository {
	userRepository := repository.NewUserRepository(mongoClient, config.MongoDB.Database, config.MongoDB.UserCollection)

	return &Repository{
		UserRepository: userRepository,
	}
}
