package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/database"
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Endast POST tillåten", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("jag kom hit")

	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ogiltig JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Mottaget username: %s, password: %s\n", req.Username, req.Password)

	// check with the database of the username exists
	var user models.User
	result := database.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "Felaktigt användarnamn eller lösenord", http.StatusUnauthorized)
		return
	}

	fmt.Println(user.Username)

	if checkCredentials(req, user) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": "dummy123",
		})
	} else {
		http.Error(w, "Felaktigt användarnamn eller lösenord", http.StatusUnauthorized)
	}
}

func checkCredentials(req loginRequest, userDB models.User) bool {
	var err error = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(req.Password))

	return err == nil
}