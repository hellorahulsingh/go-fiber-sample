package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// CreateUserRequest defines the structure of the request body for creating a user
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// CreateUserResponse defines the structure of the response after creating a user
type CreateUserResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// LoginRequest defines the structure of the request body for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserHandler handles the creation of a new user
func CreateUserHandler(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Hash the password before saving the user
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := &User{
		ID:       primitive.NewObjectID(),
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Create the user in the database
	if err := CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(CreateUserResponse{
		ID:      user.ID.Hex(),
		Email:   user.Email,
		Message: "User created successfully",
	})
}

// GetAllUsersHandler handles GET /users to fetch all users
func GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(users)
}

// LoginHandler handles POST /auth/login for user authentication
func LoginHandler(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Fetch the user by email
	user, err := GetUserByEmail(req.Email)
	if err != nil {
		log.Println("Error fetching user:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare the entered password with the stored hash
	if !ComparePasswords(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// If password is correct, proceed with further login logic (e.g., JWT)
	return c.JSON(fiber.Map{"message": "Login successful"})
}
