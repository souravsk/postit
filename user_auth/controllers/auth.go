package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/souravsk/go-zero-to-hero/user_auth/models"
	"github.com/souravsk/go-zero-to-hero/user_auth/utils"
)

var jwtkey = []byte("My_secret_key")

// ##################################################--LOGIN--######################################################

func Login(c *gin.Context) {

	// **********************Parse Incoming JSON*******************************
	var user models.User
	// attempts to bind the JSON payload from the request body to the user variable.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// **********************Check if User Exists*******************************

	var existingUser models.User

	//  queries the database for a user with the given email.
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
	}

	// **********************Validate Password*******************************

	// The purpose of this function is to compare a plain-text password (user.Password) with a hashed password stored in the database
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	// **********************Generate JWT Token*******************************

	//sets the token's expiration time to 5 minutes from the current time.
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		User: existingUser.User,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// reates a new JWT token using the HS256 signing method and the specified claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signs the token with a secret key (jwtKey)
	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		c.JSON(500, gin.H{"error": "cloud not generate token"})
		return
	}

	// **********************Set the Token in a Cookie and Send Response*******************************

	// ets a cookie named token with the JWT token. The cookie is valid for the duration specified by expirationTime and is set for the localhost domain. The false parameter means the cookie is not marked as Secure (i.e., not only sent over HTTPS), and true marks it as HttpOnly (not accessible via JavaScript).
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "locahost", false, true)

	c.JSON(200, gin.H{"success": "user logged in"})
}

// ##################################################--Signup--######################################################

func Signup(c *gin.Context) {

	// **********************Parse Incoming JSON*******************************
	var user models.User
	// attempts to bind the JSON payload from the request body to the user variable.
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received signup data:", user)

	// **********************Check if User Exists*******************************

	var existingUser models.User
	println(existingUser.ID, "existing user ID")
	//  queries the database for a user with the given email.
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.Email == user.Email {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	// **********************Hash the Password*******************************

	var errHash error
	// the user’s password using a utility function GenerateHashPassword from the utils package.  typically returns the hashed password and any error encountered
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "cloud not generate password hasd"})
		return
	}

	// inserts the new user record into the database.
	models.DB.Create(&user)

	c.JSON(200, gin.H{"success": "user created"})

}

// ##################################################--Home--######################################################

func Home(c *gin.Context) {

	// **********************Hash the Password*******************************

	// c.Cookie("token") attempts to retrieve a cookie named "token" from the request. This cookie is expected to contain the JWT (JSON Web Token).cookie will hold the value of the cookie
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// **********************Parse the Token*******************************

	// function call that parses the JWT token to extract the claims. The cookie contains the JWT token to be parsed.
	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// **********************Check User User*******************************

	// checks if the User contained in the claims is either "user" or "admin".If the User is not "user" or "admin", it returns a 401 Unauthorized response.
	if claims.User != "user" && claims.User != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "User": claims.User})
}

// ##################################################--Premium--######################################################

func Premium(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
	}

	if claims.User != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "premium page", "User": claims.User})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
