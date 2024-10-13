//This model will contain information about the claims associated with a JSON Web Token (JWT). It will include the standard claims defined by the JWT specification, as well as a field for the user's User within the application.

package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	User string `json:"User"`
	jwt.StandardClaims
}
