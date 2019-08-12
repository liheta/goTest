package router

import (
	"demo/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	// Creates a default gin router
	r := gin.Default() // Grouping routes
	// group： v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", handlers.Login)
		v1.GET("/getUsers", handlers.GetUsers)
		v1.GET("/helloWord2", handlers.HelloPage)
		v1.GET("/helloWord", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "helloword",
			})
		})
		v1.POST("/updateUser", handlers.UpdateUser)
		v1.POST("/insertUser", handlers.InsertUser)
		v1.POST("/deleteUser", handlers.DeleteUser)
	}
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
