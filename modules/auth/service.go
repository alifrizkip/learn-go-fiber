package auth

import (
	"errors"
	"learn-go-fiber/config"
	"learn-go-fiber/modules/user"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// IService interface
type IService interface {
	GenerateToken(userID int) (string, error)
	RegisterUser(input *registerUserInput) (*user.User, error)
	Login(input *loginInput) (*user.User, error)
}

type service struct {
	userRepo user.IRepository
}

// JwtSecretKey jwt
var JwtSecretKey = config.JwtSecretKey

// NewService auth service
func NewService(userRepo user.IRepository) IService {
	return &service{userRepo}
}

func (s *service) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(config.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *service) RegisterUser(input *registerUserInput) (*user.User, error) {
	user := user.User{
		Name:  input.Name,
		Email: input.Email,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(passwordHash)

	newUser, err := s.userRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *service) Login(input *loginInput) (*user.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepo.FindBy(&user.User{Email: email})
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}