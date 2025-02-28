package controllers

import (
	"encoding/json"
	// "fmt"
	"golang/config"
	"golang/models"
	"golang/utils"
	"net/http"
)



func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Check if user already exists
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Insert new user
	_, err = config.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}


func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	row := config.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", user.Email)
	var hashedPassword string
	err := row.Scan(&user.ID, &hashedPassword)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(user.Password, hashedPassword) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.Email)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}



// UpdateUser updates user details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash new password if provided
	if user.Password != "" {
		hashedPassword, _ := utils.HashPassword(user.Password)
		user.Password = hashedPassword
	}

	_, err := config.DB.Exec("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?", 
		user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
