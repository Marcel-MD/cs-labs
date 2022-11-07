package passwords

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, password string) error
	Login(email, password string) error
}

type userService struct {
	db Database
}

func NewUserService(database Database) UserService {
	return &userService{db: database}
}

func (s *userService) Register(email, password string) error {

	_, err := s.db.Get(email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.db.Set(email, []byte(hashedPassword))
}

func (s *userService) Login(email, password string) error {
	hashedPassword, err := s.db.Get(email)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
