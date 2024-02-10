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
	r.Use(CORSMiddleware())
	//r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	public := r.Group("/api")
	public.GET("/", version)

	initializePaths(public, routes.Availbility)
	initializePaths(public, routes.Events)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, hx-request, hx-current-url")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
