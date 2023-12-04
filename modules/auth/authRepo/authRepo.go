package authrepo

import "go.mongodb.org/mongo-driver/mongo"

type (
	AuthRepoitoryService interface{}

	authRepoitory struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepoitoryService {
	return &authRepoitory{db: db}
}
