package services

import (
	"context"
	"fmt"
	"os"
	"project/repos"
	"time"

	"project/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateNewUser(context.Context, *types.User) error             // Create user
	Authenticate(context.Context, string, string) (string, error) // Login
}

type userServiceImpl struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// Create a new user with POST /users
func (u *userServiceImpl) CreateNewUser(c context.Context, user *types.User) error {

	user.Id = uuid.New().String()

	// Hashed Password Creation
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.repo.CreateUser(c, user)
}

// Create a User Session with POST /login
func (u *userServiceImpl) Authenticate(c context.Context, username, password string) (string, error) {

	// Username Verification
	user, err := u.repo.GetUserByUsername(c, username)
	if err != nil {
		return "", fmt.Errorf("errorUserNotFound")
	}

	// Password Verification
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err != nil {
		return "", fmt.Errorf("errorPassword")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Token Expiration date (+2h)
	expirationTime := time.Now().Add(2 * time.Hour)
	claims["id"] = user.Id
	claims["name"] = user.Username
	claims["exp"] = expirationTime.Unix()

	// Token Creation
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("errorToken")
	}

	return tokenString, nil
}
