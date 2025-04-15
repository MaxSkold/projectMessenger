package auth

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"regexp"
	"time"
)

type CredsInput struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Credentials struct {
	UserID      string    `gorm:"column:user_id;primaryKey"`
	Email       string    `validate:"required,email" gorm:"column:email;unique"`
	PhoneNumber string    `validate:"omitempty,e164" gorm:"column:phone_number;unique"`
	Password    string    `validate:"required,min=8" gorm:"column:passhash"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (*Credentials) TableName() string {
	return "auth.credentials"
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
		UserID:      uuid.New().String(),
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	if err := validate.Struct(creds); err != nil {
		return nil, err
	}

	var errHash error
	if creds.Password, errHash = HashPassword(password); errHash != nil {
		return nil, errHash
	}

	return creds, nil
}

func (creds *Credentials) MarshalJSON() ([]byte, error) {
	type safeCreds struct {
		UserID      string `json:"user_id"`
		Email       string `json:"Email"`
		PhoneNumber string `json:"phone_number"`
	}
	return json.Marshal(&safeCreds{
		UserID:      creds.UserID,
		Email:       creds.Email,
		PhoneNumber: creds.PhoneNumber,
	})
}

// Getters

func (creds *Credentials) GetUserID() string {
	return creds.UserID
}
func (creds *Credentials) GetEmail() string {
	return creds.Email
}
func (creds *Credentials) GetPhoneNumber() string {
	return creds.PhoneNumber
}
func (creds *Credentials) GetPassword() string {
	return creds.Password
}
