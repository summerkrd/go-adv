package auth

import (
	"encoding/json"
	"fmt"
	"go-adv/3-validation-api/verify"
	"go-adv/4-order-api-start/db"
	"net/http"
)

type AuthHandler struct {
	repo *AuthRepository
}

func NewAuthHandler(database *db.Db) *AuthHandler {
	return &AuthHandler{
		repo: NewAuthRepository(database),
	}
}

// POST /auth/send-code
func (h *AuthHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req SendCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := verify.IsValid(req); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	session, err := h.repo.CreateSession(req.Phone)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// TODO: Здесь должна быть отправка SMS через внешний сервис
	// Пока выводим код в консоль для тестирования
	fmt.Printf("SMS Code for %s: %d\n", req.Phone, session.Code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SendCodeResponse{
		SessionID: session.SessionID,
	})
}

// POST /auth/verify-code
func (h *AuthHandler) VerifyCode(w http.ResponseWriter, r *http.Request) {
	var req VerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := verify.IsValid(req); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	session, err := h.repo.VerifySession(req.SessionID, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := h.repo.GetOrCreateUser(session.Phone)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(user.ID, user.Phone)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VerifyCodeResponse{
		Token: token,
	})
}
