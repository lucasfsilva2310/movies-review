package auth

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucasfsilva2310/movies-review/internal/users"
)

type AuthService struct {
	authRepository AuthRepository
	userRepository users.UserRepository
}

func NewAuthService(authRepo *AuthRepository, userRepo *users.UserRepository) *AuthService {
	return &AuthService{
		authRepository: *authRepo,
		userRepository: *userRepo,
	}
}

func (service *AuthService) Register(user users.User) error {
	_, err := service.userRepository.GetByUsername(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return service.userRepository.Create(user)
}

func (service *AuthService) Login(request LoginRequest) (string, error) {
	username, storedPassword, err := service.authRepository.GetUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(request.Password)); err != nil {
		return "", err
	}

	token, err := service.generateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) generateJWT(username string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("secret key is not set")
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
