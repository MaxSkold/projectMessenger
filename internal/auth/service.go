package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

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
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Authentication

// TODO ...

// Register or update credentials

func (s *ServiceAuth) RegisterUser(c *CredsInput) (*Credentials, error) {
	_, err := s.repo.FindByEmail(c.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	creds, err := NewCredentials(c.Email, c.PhoneNumber, c.Password)
	if err != nil {
		return nil, err
	}

	if err := s.repo.SaveCreds(creds); err != nil {
		return nil, err
	}

	return creds, nil
}
func (s *ServiceAuth) UpdatePassword(userid, oldPassword, newPassword string) error {
	creds, err := s.repo.GetCredentialsByID(userid)
	if err != nil {
		return err
	}

	if !CheckPasswordHash(oldPassword, creds.Password()) {
		return errors.New("invalid current password")
	}

	hashedNew, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	updatedCred := &Credentials{
		userID:      creds.UserID(),
		email:       creds.Email(),
		phoneNumber: creds.PhoneNumber(),
		password:    hashedNew,
	}

	return s.repo.SaveCreds(updatedCred)
}

// User have forgotten the password ,_,

// TODO ...
