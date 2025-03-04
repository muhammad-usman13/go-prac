package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/muhammad-usman13/jwt-golang/jwt_util"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var secretKey = []byte("secret-key")

func LoginHandler(c *gin.Context) {
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("The user request value %v", u)
	fmt.Println("\nUsername: ", u.Username)
	fmt.Println("\nPassword: ", u.Password)
	if u.Username == "Cheeek" && u.Password == "123456" {
		tokenString, err := jwt_util.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No username found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func ProtectedHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]
	err := jwt_util.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area"})
}

func RefreshHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]
	err := jwt_util.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing token"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting claims"})
		return
	}
	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting username"})
		return
	}
	tokenString, err = jwt_util.CreateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
