// This model will represent a user of the application and will likely include fields such as user, email, and password.
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	User     string `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
