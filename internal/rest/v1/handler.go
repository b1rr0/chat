package v1

import (
	"chat/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initChatRouters(v1)
	}
}
