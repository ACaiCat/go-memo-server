package service

import (
	"errors"
	"regexp"

	"github.com/ACaiCat/memo/internal/model"
	"github.com/ACaiCat/memo/internal/repository"
	"github.com/ACaiCat/memo/pkg/mw"
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
	Create(email string, name string, password string) (string, uint, error)
	ValidaUser(email string, password string) (string, uint, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) Create(email string, name string, password string) (string, uint, error) {
	if !isValidEmail(email) {
		return "", 0, ErrInvalidEmail
	}

	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", 0, err
	}
	if user != nil {
		return "", 0, ErrEmailUsed
	}

	user, err = u.repo.GetByName(name)
	if err != nil {
		return "", 0, err
	}
	if user != nil {
		return "", 0, ErrNameUsed
	}

	hash, err := hashPassword(password)
	if err != nil {
		return "", 0, err
	}

	user = &model.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	if err = u.repo.Create(user); err != nil {
		return "", 0, err
	}

	token, err := mw.NewJWT(user.ID)
	if err != nil {
		return "", 0, err
	}

	return token, user.ID, err
}

func (u *userService) ValidaUser(email string, password string) (string, uint, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", 0, err
	}
	if user == nil {
		return "", 0, ErrUserNotFound
	}

	if err = validaPassword(user.Password, password); err != nil {
		return "", 0, ErrPasswordError
	}

	token, err := mw.NewJWT(user.ID)
	if err != nil {
		return "", 0, err
	}

	return token, user.ID, nil

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
