package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
}
