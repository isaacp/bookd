package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/api/handlers/availability"
)

var Availbility map[string]map[string]func(c *gin.Context) = map[string]map[string]func(c *gin.Context){
	"/availability": {
		"GET": availability.List,
	},
}
