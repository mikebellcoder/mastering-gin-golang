package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(200, "Hello, %s!", name)
	})

	r.GET("/search", func(ctx *gin.Context) {
		query := ctx.DefaultQuery("q", "none")
		ctx.String(200, "Searched for: %s", query)
	})

	r.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name": "John",
			"age":  25,
		})
	})

	api := r.Group("/api")

	tasks := []gin.H{
		{
			"name": "John",
			"age":  25,
		},
		{
			"name": "John",
			"age":  25,
		},
		{
			"name": "John",
			"age":  25,
		},
	}

	{
		api.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(200, tasks)
		})

		api.POST("/tasks", func(ctx *gin.Context) {
			tasks = append(tasks, gin.H{
				"name": "New",
				"age": 1,
			})
			ctx.String(200, "There are now %d tasks", len(tasks))
		})
	}

	r.Run(":8080")
}
