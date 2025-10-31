package handlers

import (
	"net/http"

	"github.com/dedenurr/contactflow/api-contact-form/responses"

	"github.com/gin-gonic/gin"
)

type MainHandler struct{}

func NewMainHandler() *MainHandler{
	return &MainHandler{}
}

func (h *MainHandler) MainHandler(c *gin.Context){
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Welcome to the API Contact Form Service.",
	})
}