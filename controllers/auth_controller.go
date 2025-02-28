package controllers

import (
	"encoding/json"
	"fmt"
	"golang/config"
	"golang/models"
	"golang/utils"
	"net/http"
)


func Register(w http.ResponseWriter, r *http.Request) {

	fmt.Println("in Controller...")

	var registration models.Registration

	json.NewDecoder(r.Body).Decode(&registration)


	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", registration.Email).Scan(&count)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if count > 0 {

		json.NewEncoder(w).Encode(map[string]string{"status":"0","message": "User already exists"})
		return
	}

	hashedPassword, _ := utils.HashPassword(registration.Password)
	registration.Password = hashedPassword

	_, err = config.DB.Exec("INSERT INTO users (name, email, password,address,gender,pincode,date_of_birth,city,phone,state,status,entered_by,updated_by,user_category,entered_date_time,updated_date_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",registration.Name, registration.Email, registration.Password,registration.Address,registration.Gender,registration.Pincode,registration.DateOfBirth,registration.City,registration.Phone,registration.State,registration.Status,"1","1","owner","now()","now()")

	// fmt.Println("INSERT INTO users (name, email, password,address,gender,pincode,date_of_birth,city,phone,state,status,entered_by,updated_by,user_category,entered_date_time,updated_date_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", registration.Name, registration.Email, registration.Password,registration.Address,registration.Gender,registration.Pincode,registration.DateOfBirth,registration.City,registration.Phone,registration.State,registration.Status,"1","1","owner","now()","now()")
	
	if err != nil {

		json.NewEncoder(w).Encode(map[string]string{"status":"0","message": "failed user creation"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status":"1","message": "User registered successfully"})
}


func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var hashedPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Email).Scan(&hashedPassword)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "0", "msg": "Email or Password Incorrect"})
		return
	}

	if !utils.CheckPasswordHash(user.Password, hashedPassword) {
		json.NewEncoder(w).Encode(map[string]string{"status": "0", "msg": "Email or Password Incorrect"})
		return
	}

	token, _ := utils.GenerateJWT(user.Email)
	json.NewEncoder(w).Encode(map[string]string{"token": token, "status": "1", "msg": "Login Successfully"})
}

