package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/api/handlers/events"
)

var Events map[string]map[string]func(c *gin.Context) = map[string]map[string]func(c *gin.Context){
	"/events": {
		"GET": events.List,
	},
}
