//go:generate mockery --name=IUserRepository
package repository

import (
	"context"

	"github.com/PongponZ/go-book-store/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	DeleteByID(ctx context.Context, id string) error
}

type userRepository struct {
	db             *mongo.Client
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Client, databaseName string, userCollectionName string) IUserRepository {
	userCollection := db.Database(databaseName).Collection(userCollectionName)

	return &userRepository{db: db, userCollection: userCollection}
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	filter := bson.M{"_id": id}

	var user entity.User

	err := r.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *userRepository) DeleteByID(ctx context.Context, id string) error {
	return nil
}
