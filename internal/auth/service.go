package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrAlreadyExists = errors.New("user already exists")

type ServiceAuth struct {
	repo CredRepo
}

func NewServiceAuth(repo CredRepo) *ServiceAuth {
	return &ServiceAuth{repo}
}

// Common functions

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
func IsValid(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// GenerateConfirmationCode function for 2FA confirmation code generate
func GenerateConfirmationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(100000))
}

// Authentication

func (s *ServiceAuth) LoginUser(c *CredsInput) error {
	var data *Credentials
	var err error

	switch {
	case c.PhoneNumber != "":
		data, err = s.repo.FindByPhoneNumber(c.PhoneNumber)
	case c.Email != "":
		data, err = s.repo.FindByEmail(c.Email)
	default:
		return errors.New("either phone number or email must be provided")
	}

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("creds not found")
	}

	if !IsValid(data.Password, c.Password) {
		return errors.New("password is invalid")
	}

	return nil
}

// Register or update credentials

func (s *ServiceAuth) RegisterUser(c *CredsInput) error {
	_, err := s.repo.FindByEmail(c.Email)
	if err == nil {
		return ErrAlreadyExists
	}
	if !errors.Is(err, ErrNotFound) {
		return err
	}

	creds, err := NewCredentials(c.Email, c.PhoneNumber, c.Password)
	if err != nil {
		return err
	}

	if err := s.repo.SaveCreds(creds); err != nil {
		return err
	}

	return nil
}
func (s *ServiceAuth) UpdatePassword(userid, oldPassword, newPassword string) error {
	creds, err := s.repo.GetCredentialsByID(userid)
	if err != nil {
		return err
	}

	if !IsValid(oldPassword, creds.GetPassword()) {
		return errors.New("invalid current Password")
	}

	hashedNew, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	updatedCred := &Credentials{
		UserID:      creds.GetUserID(),
		Email:       creds.GetEmail(),
		PhoneNumber: creds.GetPhoneNumber(),
		Password:    hashedNew,
	}

	return s.repo.SaveCreds(updatedCred)
}

func (s *ServiceAuth) ConfirmEmail(code string) error {
	creds, err := s.repo.FindByConfirmationCode(code)
	if err != nil {
		return err
	}

	if creds.CodeExpiresAt == nil || time.Now().After(*creds.CodeExpiresAt) {
		return errors.New("confirmation code expired")
	}

	creds.EmailConfirmed = true
	creds.ConfirmationCode = nil
	creds.CodeExpiresAt = nil

	return s.repo.SaveCreds(creds)
}

// User have forgotten the Password ,_,

// TODO ...
