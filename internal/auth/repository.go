package auth

import (
	"database/sql"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("credentials not found")

type (
	CredRepo interface {
		GetCredentialsByID(id string) (*Credentials, error)
		FindByEmail(email string) (*Credentials, error)
		SaveCreds(cred *Credentials) error
		RemoveCreds(id string) error
	}

	PostgresCredRepo struct {
		mu sync.RWMutex
		db *sql.DB
	}
	MapsCredRepo struct {
		mu   sync.RWMutex
		repo map[string]*Credentials
	}
)

// ------------------ Work with PSQL database ----------------------

func NewPostgresCredRepo(db *sql.DB) *PostgresCredRepo {
	return &PostgresCredRepo{
		db: db,
	}
}

func (repo *PostgresCredRepo) GetCredentialsByID(id string) (*Credentials, error) {
	// TODO: реализация через INSERT INTO ...
	return nil, nil
}
func (repo *PostgresCredRepo) SaveCreds(cred *Credentials) error {
	// TODO: реализация через SELECT ... WHERE credential_id = $1
	return nil
}
func (repo *PostgresCredRepo) FindByEmail(email string) (*Credentials, error) {
	// TODO ...
	return nil, nil
}

// ----------------- Work with Map (in-memory) ---------------------

func NewMapsCredRepo() *MapsCredRepo {
	return &MapsCredRepo{
		repo: make(map[string]*Credentials, 10),
	}
}

func (repo *MapsCredRepo) GetCredentialsByID(id string) (*Credentials, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	cred, exists := repo.repo[id]
	if !exists {
		return nil, ErrNotFound
	}
	return cred, nil
}
func (repo *MapsCredRepo) SaveCreds(cred *Credentials) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.repo[cred.UserID()] = cred
	return nil
}
func (repo *MapsCredRepo) FindByEmail(email string) (*Credentials, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, cred := range repo.repo {
		if cred.Email() == email {
			return cred, nil
		}
	}
	return nil, ErrNotFound
}
func (repo *MapsCredRepo) RemoveCreds(userID string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.repo[userID]; !ok {
		return errors.New("user not found")
	}

	delete(repo.repo, userID)
	return nil
}
