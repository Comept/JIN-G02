package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getSubMsgById(c *gin.Context) {
	id,_ := c.Get(userCtx)
	id_sub:= id.(int)
	id_pub, err := strconv.Atoi(c.Param("pub_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	messages, err := h.services.Publisher.ShowById(id_sub, id_pub)
	if err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messages)
	return
}

func (h *Handler) getSubMsg(c *gin.Context) {
	id,_ := c.Get(userCtx)
	messages, err := h.services.Publisher.ShowAll(id.(int))
	if err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messages)
	return
}
