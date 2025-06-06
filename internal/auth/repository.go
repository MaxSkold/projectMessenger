package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/smtp"
	"sync"
)

var ErrNotFound = errors.New("credentials not found")

type (
	CredRepo interface {
		GetCredentialsByID(id string) (*Credentials, error)
		FindByEmail(email string) (*Credentials, error)
		FindByPhoneNumber(phoneNumber string) (*Credentials, error)
		SaveCreds(cred *Credentials) error
		RemoveCreds(id string) error
		FindByConfirmationCode(confirmCode string) (*Credentials, error)
	}

	Mailer interface {
		SendConfirmationEmail(emailTo, code string) error
	}

	PostgresCredRepo struct {
		mu sync.RWMutex
		db *gorm.DB
	}
	MapsCredRepo struct {
		mu   sync.RWMutex
		repo map[string]*Credentials
	}

	SMTPMailer struct {
		From     string
		Password string
		Host     string
		Port     int
	}
)

// ------------------ Work with PSQL database ----------------------

func NewPostgresCredRepo(db *gorm.DB) *PostgresCredRepo {
	return &PostgresCredRepo{
		db: db,
	}
}

func (psql *PostgresCredRepo) GetCredentialsByID(id string) (*Credentials, error) {
	psql.mu.RLock()
	defer psql.mu.RUnlock()

	var cred Credentials
	if err := psql.db.Where("user_id = ?", id).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &cred, nil
}
func (psql *PostgresCredRepo) SaveCreds(cred *Credentials) error {
	psql.mu.Lock()
	defer psql.mu.Unlock()

	return psql.db.Save(cred).Error
}
func (psql *PostgresCredRepo) FindByEmail(email string) (*Credentials, error) {
	psql.mu.RLock()
	defer psql.mu.RUnlock()

	var cred Credentials
	if err := psql.db.Where("email = ?", email).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &cred, nil
}
func (psql *PostgresCredRepo) FindByPhoneNumber(phoneNumber string) (*Credentials, error) {
	psql.mu.RLock()
	defer psql.mu.RUnlock()

	var cred Credentials
	if err := psql.db.Where("phone_number = ?", phoneNumber).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &cred, nil
}
func (psql *PostgresCredRepo) FindByConfirmationCode(confirmCode string) (*Credentials, error) {
	psql.mu.RLock()
	defer psql.mu.RUnlock()

	var cred Credentials
	if err := psql.db.Where("confirmation_code = ?", confirmCode).First(&cred).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &cred, nil
}
func (psql *PostgresCredRepo) RemoveCreds(id string) error {
	psql.mu.Lock()
	defer psql.mu.Unlock()

	return psql.db.Delete(&Credentials{}, id).Error
}

// ----------------- Work with Map (in-memory) ---------------------

//func NewMapsCredRepo() *MapsCredRepo {
//	return &MapsCredRepo{
//		repo: make(map[string]*Credentials, 10),
//	}
//}
//
//func (repo *MapsCredRepo) GetCredentialsByID(id string) (*Credentials, error) {
//	repo.mu.RLock()
//	defer repo.mu.RUnlock()
//
//	cred, exists := repo.repo[id]
//	if !exists {
//		return nil, ErrNotFound
//	}
//	return cred, nil
//}
//func (repo *MapsCredRepo) SaveCreds(cred *Credentials) error {
//	repo.mu.Lock()
//	defer repo.mu.Unlock()
//
//	repo.repo[cred.GetUserID()] = cred
//	return nil
//}
//func (repo *MapsCredRepo) FindByEmail(email string) (*Credentials, error) {
//	repo.mu.RLock()
//	defer repo.mu.RUnlock()
//
//	for _, cred := range repo.repo {
//		if cred.GetEmail() == email {
//			return cred, nil
//		}
//	}
//	return nil, ErrNotFound
//}
//func (repo *MapsCredRepo) RemoveCreds(userID string) error {
//	repo.mu.Lock()
//	defer repo.mu.Unlock()
//
//	if _, ok := repo.repo[userID]; !ok {
//		return errors.New("user not found")
//	}
//
//	delete(repo.repo, userID)
//	return nil
//}

// ----------------- Work with eMail -------------------------------

func NewSMTPMailer() *SMTPMailer {
	return &SMTPMailer{
		From:     "mrk642004@gmail.com",
		Password: "napmi9-sYxfos-darvuf",
		Host:     "smtp.gmail.com",
		Port:     587,
	}
}

func (psm *SMTPMailer) SendConfirmationEmail(emailTo, code string) error {
	auth := smtp.PlainAuth("", psm.From, psm.Password, psm.Host)

	subject := "Confirm your emailTo"
	body := fmt.Sprintf("Hello!\n\nPlease confirm your emailTo using this code:\n\n%s\n\nThank you!", code)

	msg := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s",
		emailTo, subject, body,
	))

	addr := fmt.Sprintf("%s:%d", psm.Host, psm.Port)

	return smtp.SendMail(addr, auth, emailTo, []string{emailTo}, msg)
}
