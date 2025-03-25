package handlers

import (
	"encoding/json"
	"net/http"

	"backend/database"
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

// Separat input-struct för registrering
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Separat output-struct (respons till klienten)
type RegisterResponse struct {
	Token string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Endast POST tillåten", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ogiltig JSON", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Kunde inte kryptera lösenord", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Användarnamn finns redan", http.StatusConflict)
		return
	}

	// Returnera JSON med token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": "dummy123",
	})
}

