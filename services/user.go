package services

import (
	"context"
	"os"

	"../repos"
	"../types"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateNewUser(context.Context, *types.User) error
	GetUser(context.Context, string) (*types.User, error)
	Authenticate(context.Context, string, string) (string, error) // New method for authentication
}

type userServiceImpl struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

/*
TODO: Add service that ENCRYPTS PASSWORD and STORES IN DB

Tips
- Use bcrypt package!
- Save bcrypt secret on .env and load it in App configuration!
- Inject app configuration (bcrypt secret) into here (user service)
*/
func (u *userServiceImpl) CreateNewUser(c context.Context, user *types.User) error {
	// Generate a new UUID for the user's ID
	user.Id = uuid.New().String()

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Store the user in the repository
	return u.repo.CreateUser(c, user)
}

func (u *userServiceImpl) Authenticate(c context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(c, username)
	if err != nil {
		return "", err
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["name"] = user.Username
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
