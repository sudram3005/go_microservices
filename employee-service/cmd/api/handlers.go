package main

import (
	"employee/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) GetAllUser(ctx *gin.Context) {
	requestPayload, err := app.Models.User.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = ctx.BindJSON(&requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all users",
		Data:    requestPayload,
	}

	ctx.JSON(http.StatusAccepted, payload)
}

func (app *Config) GetOne(ctx *gin.Context) {
	var requestPayload data.User

	err := ctx.BindJSON(&requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	id := requestPayload.ID
	data, err := app.Models.User.GetOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all users",
		Data:    data,
	}

	ctx.JSON(http.StatusAccepted, payload)
}

func (app *Config) DeleteUser(ctx *gin.Context) {
	var requestPayload data.User

	err := ctx.BindJSON(&requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = requestPayload.DeleteByID(requestPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "User deleted",
	}

	ctx.JSON(http.StatusAccepted, payload)
}

func (app *Config) AddUser(ctx *gin.Context) {
	var requestPayload data.User

	err := ctx.BindJSON(&requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := requestPayload.Insert(requestPayload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "User deleted",
		Data:    data,
	}

	ctx.JSON(http.StatusAccepted, payload)
}
