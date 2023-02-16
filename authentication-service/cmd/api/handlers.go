package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Authenticate(ctx *gin.Context) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.BindJSON(&requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	//validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	ctx.JSON(http.StatusAccepted, payload)
}
