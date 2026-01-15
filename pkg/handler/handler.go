package handler

import (
	"TODOAPP/pkg/service"

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
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllLists)
			lists.GET("/:id", h.GetAllListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetAllItem)
				items.GET("/:item_id", h.GetAllItemById)
				items.PUT("/:item_id", h.UpdateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}

	return router
}
