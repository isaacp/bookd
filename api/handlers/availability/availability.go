package availability

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/api/core/entities"
)

func List(c *gin.Context) {
	manager := entities.CalendarManager{}
	freeIntervals := manager.FreeIntervals(c.Query("start"), c.Query("end"))

	c.JSON(http.StatusOK, freeIntervals)
}
