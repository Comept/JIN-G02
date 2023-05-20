package handler

import (
	"sth/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		massege := api.Group("/message")
		{
			pub := massege.Group("/pub")
			{
				pub.POST("/newMes", h.createMessage)
				pub.GET("/:sub_id", h.getPubMsgById)
				pub.GET("/all", h.getPubMsg)
			}
			sub := massege.Group("/sub")
			{
				sub.GET("/:pub_id", h.getSubMsgById)
				sub.GET("/all", h.getSubMsg)
			}
		}
	}
	return router
}
