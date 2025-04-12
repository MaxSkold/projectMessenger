package auth

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"regexp"
)

type CredsInput struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Credentials struct {
	userID      string // Первичный ключ. id как паспорт в общем
	email       string `validate:"required,email"`
	phoneNumber string `validate:"omitempty,e164"`
	password    string `validate:"required,min=8"`
}

// NewCredentials function create credentials for user
func NewCredentials(email, phoneNumber, password string) (*Credentials, error) {
	validate := validator.New()

	_ = validate.RegisterValidation("e164", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		e164Pattern := `^\+?[1-9]\d{1,14}$`
		regex := regexp.MustCompile(e164Pattern)
		return regex.MatchString(phone)
	})

	creds := &Credentials{
		userID:      uuid.New().String(),
		email:       email,
		phoneNumber: phoneNumber,
		password:    password,
	}

	if err := validate.Struct(creds); err != nil {
		return nil, err
	}

	var errHash error
	if creds.password, errHash = HashPassword(password); errHash != nil {
		return nil, errHash
	}

	return creds, nil
}

func (creds *Credentials) MarshalJSON() ([]byte, error) {
	type safeCreds struct {
		UserID      string `json:"user_id"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
	return json.Marshal(&safeCreds{
		UserID:      creds.userID,
		Email:       creds.email,
		PhoneNumber: creds.phoneNumber,
	})
}

// Getters

func (creds *Credentials) UserID() string {
	return creds.userID
}
func (creds *Credentials) Email() string {
	return creds.email
}
func (creds *Credentials) PhoneNumber() string {
	return creds.phoneNumber
}
func (creds *Credentials) Password() string {
	return creds.password
}
