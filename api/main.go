package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/bookd/routes"
)

func main() {
	os.Setenv("PORT", "8080")

	r := gin.Default()
	public := r.Group("/api")

	public.GET("/", version)

	initializePaths(public, routes.Availbility)
	initializePaths(public, routes.Events)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initializePaths(group *gin.RouterGroup, routes map[string]map[string]func(c *gin.Context)) {
	for path, content := range routes {
		for verb, function := range content {
			switch verb {
			case "DELETE":
				group.DELETE(path, function)
			case "GET":
				group.GET(path, function)
			case "PATCH":
				group.PATCH(path, function)
			case "POST":
				group.POST(path, function)
			case "PUT":
				group.PUT(path, function)
			}
		}
	}
}

func version(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, `v0.1`)
}
