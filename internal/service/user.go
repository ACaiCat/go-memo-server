package service

import (
	"errors"
	"regexp"

	"github.com/ACaiCat/memo/internal/model"
	"github.com/ACaiCat/memo/internal/mw"
	"github.com/ACaiCat/memo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound  = errors.New("service: user not found")
	ErrEmailUsed     = errors.New("service: email already in use")
	ErrInvalidEmail  = errors.New("service: invalid email format")
	ErrNameUsed      = errors.New("service: name already in use")
	ErrPasswordError = errors.New("service: password error")
)

type UserService interface {
	Create(email string, name string, password string) (string, error)
	ValidaUser(email string, password string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) Create(email string, name string, password string) (string, error) {
	if !isValidEmail(email) {
		return "", ErrInvalidEmail
	}

	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user != nil {
		return "", ErrEmailUsed
	}

	user, err = u.repo.GetByName(name)
	if err != nil {
		return "", err
	}
	if user != nil {
		return "", ErrNameUsed
	}

	hash, err := hashPassword(password)
	if err != nil {
		return "", err
	}

	user = &model.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	if err = u.repo.Create(user); err != nil {
		return "", err
	}

	token, err := mw.NewJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *userService) ValidaUser(email string, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrUserNotFound
	}

	if err = validaPassword(user.Password, password); err != nil {
		return "", ErrPasswordError
	}

	token, err := mw.NewJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func validaPassword(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}
