package user

import (
	"context"
	"go-fiber-app/internal/config"
	"golang.org/x/crypto/bcrypt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func getUserCollection() *mongo.Collection {
	return config.MongoDB.Collection("go-users")
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares the entered password with the stored hash
func ComparePasswords(storedPassword, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword))
	return err == nil
}

// CreateUser creates a new user and stores it in the database
func CreateUser(user *User) error {
	// Hash the password before storing it
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}
	user.Password = hashedPassword

	// Insert user into the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = getUserCollection().InsertOne(ctx, user)
	return err
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := getUserCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []User
	err = cursor.All(ctx, &users)
	return users, err
}

// GetUserByEmail fetches a user by their email
func GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user User
	err := getUserCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
