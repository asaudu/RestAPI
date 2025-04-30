package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"addyCodes.com/RestAPI/models"
	"addyCodes.com/RestAPI/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context, db *sql.DB) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save(db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context, db *sql.DB) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	fmt.Printf("User check %v", user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials(db)
	fmt.Println("User ID Check 2", user.ID)
	fmt.Printf("validate method check %v", err)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
