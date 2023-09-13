package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saktialfansyahp/go-rest-api/config"
	"github.com/saktialfansyahp/go-rest-api/helper"
	"github.com/saktialfansyahp/go-rest-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var userInput models.AuthRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).Preload("Role").First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"message": "username atau password salah"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "username atau password salah"})
		return
	}

	expTime := time.Now().Add(time.Minute * 30)
	claims := &config.JWTClaim{
		Username:         user.Username,
		Role:             user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "data": user, "token": token})
}

// func Register(w http.ResponseWriter, r *http.Request) {
// 	var userInput models.AuthRequest
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&userInput); err != nil {
// 		response := map[string]string{"message": err.Error()}
// 		helper.ResponseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}
// 	defer r.Body.Close()

// 	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
// 	userInput.Password = string(hashPassword)

// 	var role models.Role
// 	if err := models.DB.First(&role, userInput.RoleID).Error; err != nil {
// 		response := map[string]string{"message": err.Error()}
// 		helper.ResponseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	user := models.User{
// 		Name: userInput.Name,
// 		Username: userInput.Username,
// 		Password: userInput.Password,
// 		RoleID: userInput.RoleID,
// 		Role: role,
// 	}

// 	if  err := models.DB.Create(&user).Error; err != nil {
// 		response := map[string]string{"message": err.Error()}
// 		helper.ResponseJSON(w, http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := map[string]interface{}{"message": "success", "data": user}
// 	helper.ResponseJSON(w, http.StatusOK, response)

// }

func Register(c *gin.Context) {
	var userInput models.AuthRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	var role models.Role
	if err := models.DB.First(&role, userInput.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := models.User{
		Name:     userInput.Name,
		Username: userInput.Username,
		Password: userInput.Password,
		RoleID:   userInput.RoleID,
		Role:     role,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
	})

	response := map[string]string{"message": "logout success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Role(w http.ResponseWriter, r *http.Request) {
	var userInput models.Role

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if  err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "success", "data": userInput}
	helper.ResponseJSON(w, http.StatusOK, response)
}