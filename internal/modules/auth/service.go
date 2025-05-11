package auth

import (
	"context"
	"go-fiber-app/internal/config"
	userModule "go-fiber-app/internal/modules/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Authenticate(email, password string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := config.MongoDB.Collection("go-users")

	var user userModule.User
	err := coll.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return false, err
	}

	return true, nil
}
