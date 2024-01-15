package v1

import (
	"chat/domain"
	"chat/internal/rest/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetTasks godoc
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param taskname query string false "Taskname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []tasks.Task
// @Router /api/tasks [get]
// @Security Authorization Token
func (h *Handler) getChatMessages(c *gin.Context) {
	idd, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	messages, err := h.services.Message.FindByTo(c, idd)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// GetTasks godoc
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param taskname query string false "Taskname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []tasks.Task
// @Router /api/tasks [get]
// @Security Authorization Token
func (h *Handler) getOneToOneMessages(c *gin.Context) {
	var inputData models.FromToModel

	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Message.FindMessagesBetweenUsers(c, inputData.From, inputData.To)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

// GetTasks godoc
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param taskname query string false "Taskname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []tasks.Task
// @Router /api/tasks [get]
// @Security Authorization Token
func (h *Handler) postMessage(c *gin.Context) {
	var inputData domain.Message

	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Message.Create(c, inputData)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) initChatRouters(rGroup *gin.RouterGroup) {
	rGroup.GET("/chat/:id", h.getChatMessages)
	rGroup.GET("/chat/personal", h.getOneToOneMessages)
	rGroup.POST("/chat/messages", h.postMessage)
}
