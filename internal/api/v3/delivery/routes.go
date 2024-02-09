package delivery

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
			auth.POST("/refresh", h.Refresh)
			auth.GET("/sessions", h.GetAllSessions)
		}
	}
	return router
}
