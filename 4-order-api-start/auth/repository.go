package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"go-adv/4-order-api-start/db"
	"go-adv/4-order-api-start/models"
	"math/big"
	"time"
)

type AuthRepository struct {
	Database *db.Db
}

func NewAuthRepository(database *db.Db) *AuthRepository {
	return &AuthRepository{Database: database}
}

// Генерация случайного sessionId
func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Генерация 4-значного кода
func generateCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + 1000, nil
}

// Создание сессии
func (r *AuthRepository) CreateSession(phone string) (*models.Session, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return nil, err
	}

	code, err := generateCode()
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		SessionID: sessionID,
		Phone:     phone,
		Code:      code,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Verified:  false,
	}

	result := r.Database.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

// Верификация кода
func (r *AuthRepository) VerifySession(sessionID string, code int) (*models.Session, error) {
	var session models.Session
	result := r.Database.Where("session_id = ? AND code = ?", sessionID, code).First(&session)

	if result.Error != nil {
		return nil, errors.New("invalid session or code")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("session expired")
	}

	if session.Verified {
		return nil, errors.New("session already used")
	}

	session.Verified = true
	r.Database.Save(&session)

	return &session, nil
}

// Создание или получение пользователя
func (r *AuthRepository) GetOrCreateUser(phone string) (*models.User, error) {
	var user models.User
	result := r.Database.Where("phone = ?", phone).First(&user)

	if result.Error != nil {
		user = models.User{Phone: phone}
		if err := r.Database.Create(&user).Error; err != nil {
			return nil, err
		}
	}

	return &user, nil
}
