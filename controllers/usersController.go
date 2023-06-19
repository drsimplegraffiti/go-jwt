package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/drsimplegraffiti/gojwt/initializers"
	"github.com/drsimplegraffiti/gojwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	var body struct{
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
}

//check if user exists
var count int64
initializers.DB.Model(&models.User{}).Where("email = ?", body.Email).Count(&count)
if count > 0 {
	c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	return
}

//hash password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
	return
}

//create user
user := models.User{
	Email: body.Email,
	Password: string(hashedPassword),
}
result := initializers.DB.Create(&user)
if result.Error != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
	return
}

c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}


func Login(c *gin.Context){
	var body struct{
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
}

//check if user exists
var user models.User
initializers.DB.First(&user, "email = ?", body.Email)
if user.ID == 0 {
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	return
}

//check password
err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
if err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	return
}

//generate token
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"id": user.ID,
	"email": user.Email,
	"exp": time.Now().Add(time.Hour * 24).Unix(),
})

//sign token
tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error signing token"})
	return
}

// set cookie
c.SetSameSite(http.SameSiteNoneMode)
c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

c.JSON(http.StatusOK, gin.H{
	"token": tokenString,
	"message": "Logged in successfully",
	"email": user.Email,
	"id": user.ID,
})

}


func Logout(c *gin.Context){
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")

	// Ensure the user value is of type jwt.MapClaims
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		// Handle the error condition where the user value is not of the expected type
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user"})
		return
	}

	// Get the id from the user
	id := claims["id"].(float64)

	// Get the user from the database
	var dbUser models.User
	initializers.DB.First(&dbUser, int(id))

	c.JSON(http.StatusOK, gin.H{"user": user, "dbUser": dbUser})
}
