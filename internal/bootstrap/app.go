package bootstrap

import (
	"context"
	"time"

	"github.com/PongponZ/go-book-store/config"
	"github.com/PongponZ/go-book-store/internal/delivery/http"
	"github.com/PongponZ/go-book-store/pkg/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	repositories *Repository
	usecases     *Usecase
	mongoClient  *mongo.Client
	HTTPServer   *echo.Echo
}

func Initialize(config *config.Config) (*App, error) {
	mongoClient, err := mongodb.Connect(&mongodb.MongoConfig{
		Username: config.MongoDB.Username,
		Password: config.MongoDB.Password,
		Host:     config.MongoDB.Host,
		Port:     config.MongoDB.Port,
		Database: config.MongoDB.Database,
	})
	if err != nil {
		return nil, err
	}

	repositories := initializeRepository(config, mongoClient)
	usecases := initializeUsecase(repositories)
	handlers := initializeHTTPHandler(usecases)

	server := echo.New()
	http.RegisterRoutes(server, handlers)

	return &App{
		repositories: repositories,
		usecases:     usecases,
		mongoClient:  mongoClient,
		HTTPServer:   server,
	}, nil
}

// for graceful shutdown
func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := a.HTTPServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := a.mongoClient.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
