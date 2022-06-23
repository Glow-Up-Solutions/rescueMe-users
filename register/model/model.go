package model

import "github.com/jinzhu/gorm"

type Usuarios struct {
	gorm.Model
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Birthday   string `json:"birthday"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	City       string `json:"city"`
	Code_zip   string `json:"code_zip"`
	State      string `json:"state"`
}
