package models

type Registration struct {
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
    Address string `json:"address"`
	Gender string `json:"gender"`
	Pincode  int    `json:"id"`
	DateOfBirth string `json:"dateOfBirth"`
	City string `json:"city"`
	Phone int `json:"phone"`
	State string `json:"state"`
	Status string `json:"status"`
}

