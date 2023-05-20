package handler

import (
	"net/http"
	"sth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input sth.Users

	if err := c.BindJSON(&input); err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id,err := h.services.Authorization.CreateUser(input)
	
	if err!=nil {
		newErrorResponse(c, http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (h *Handler) signIn(c *gin.Context) {
	var input signInput
	 if err := c.BindJSON(&input); err !=nil{
		panic(err)
	}
	token, _ := h.services.Authorization.GenerateToken(input.Username, input.Password)
	c.JSON(http.StatusOK, map[string]interface{}{
	 	"token": token,
	})
	return
}