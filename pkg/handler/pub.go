package handler

import (
	"net/http"
	"sth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createMessage(c *gin.Context) {
	id,_ := c.Get(userCtx)
	var message sth.Message
	if err := c.BindJSON(&message); err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	message.Pub = id.(int)
	id, err := h.services.Message.Create(message)
	if err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
		"отправитель":message.Pub,
		"сообщение":message.Message,
		"получатель":message.Sub,
	})
	return
}

func (h *Handler) getPubMsgById(c *gin.Context) {
	id,_ := c.Get(userCtx)
	id_pub:= id.(int)
	id_sub, err := strconv.Atoi(c.Param("sub_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	messages, err := h.services.Publisher.ShowById(id_pub, id_sub)
	if err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messages)
	return
}

func (h *Handler) getPubMsg(c *gin.Context) {
	id,_ := c.Get(userCtx)
	messages, err := h.services.Publisher.ShowAll(id.(int))
	if err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messages)
	return
}