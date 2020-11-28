package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login route
func Login(c *gin.Context) {
	fmt.Println("msukn login")
	var accountData models.Account

	username := c.PostForm("username")
	password := c.PostForm("password")

	if err := config.GetDB().Where("username = ? AND password = ?", username, password).First(&accountData).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	fmt.Println(accountData)
	var jwtToken = createToken(&accountData)

	c.JSON(200, gin.H{
		"data":    accountData,
		"token":   jwtToken,
		"message": "Berhasil Login",
	})
}

func createToken(account *models.Account) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": account.ID,
		"admin_role": account.AdminRole,
		"exp":        time.Now().AddDate(0, 0, 1).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
