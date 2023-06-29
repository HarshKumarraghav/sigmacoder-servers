package auth

import (
	"errors"
	"os"
	"sigmacoder/pkg"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// The above type defines a service interface with methods for login, phone OTP login, and user sign
// up.
// @property Login - The Login method takes an email and password as parameters and returns a string
// (presumably a token or session ID) and an error. It is used to authenticate a user with their email
// and password.
// @property LoginPhoneOtp - This method is used to log in a user using their phone number and a
// one-time password (OTP). It takes the phone number as input and returns a token string and an error
// if any.
// @property SignUp - The SignUp method is used to create a new user account. It takes an input
// parameter of type InUser, which represents the user information such as email, password, and phone
// number. It returns a string representing the user ID and an error if any error occurs during the
// signup process.
type Service interface {
	Login(email string, password string) (string, error)
	LoginPhoneOtp(phone string) (string, error)
	SignUp(in InUser) (string, error)
}

// The type Svc represents a service that has a dependency on a Repo.
// @property repo - The `repo` property is a pointer to an instance of the `Repo` struct. It is used to
// access and manipulate data in the repository.
type Svc struct {
	repo *Repo
}


// The `SignUp` function is a method of the `Svc` struct that implements the `SignUp` method of the
// `Service` interface. It is responsible for handling user sign up functionality.
func (s *Svc) SignUp(in InUser) (string, error) {
	user, err := s.repo.ReadByEmail(in.Email)
	if !(err == pkg.ErrUserNotFound) && err != nil {
		return "", err
	}
	if user.Email == in.Email {
		return "", errors.New("user with id already exists")
	}
	create, err := s.repo.Create(in)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": create.ID,
		"email":  create.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return refresh, nil
}


// The `Login` function is a method of the `Svc` struct that implements the `Login` method of the
// `Service` interface. It takes an `email` and `password` as input parameters and returns a string and
// an error.
func (s *Svc) Login(email string, password string) (string, error) {
	user, err := s.repo.ReadByEmail(email)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return refresh, nil

}


// The `LoginPhoneOtp` function is a method of the `Svc` struct that implements the `LoginPhoneOtp`
// method of the `Service` interface. It takes a `phone` number as an input parameter and returns a
// string and an error.
func (s *Svc) LoginPhoneOtp(phone string) (string, error) {
	user, err := s.repo.ReadByPhoneNumber(phone)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return refresh, nil

}

// The function creates a new instance of a service with a given repository.
func NewAuthService(repo *Repo) Service {
	return &Svc{
		repo: repo,
	}
}
