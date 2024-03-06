package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/api/core/entities"
)

func List(c *gin.Context) {
	manager := entities.CalendarManager{}
	events := manager.Events(c.Query("start"), c.Query("end"))

	c.JSON(http.StatusOK, events)
}
