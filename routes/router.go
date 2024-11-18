package routes

import (
	"net/http"
	"quiz-go/controllers"
	"quiz-go/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(2 * time.Second) // Simulasi delay, bisa dihapus jika tidak diperlukan
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Sukses Menjalankan API ⚡⚡",
		})
	})
	router.POST("/login", controllers.Login)

	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware())
	{
		authorized.GET("/book", controllers.GetAllBooks)
		authorized.POST("/book", controllers.InsertBook)
		authorized.PUT("/book/:id", controllers.UpdateBook)
		authorized.DELETE("/book/:id", controllers.DeleteBook)

		authorized.GET("/category", controllers.GetAllCategories)
		authorized.POST("/category", controllers.InsertCategory)
		authorized.PUT("/category/:id", controllers.UpdateCategory)
		authorized.DELETE("/category/:id", controllers.DeleteCategory)

		authorized.GET("/user", controllers.GetAllUsers)
		authorized.POST("/user", controllers.InsertUser)
		authorized.PUT("/user/:id", controllers.UpdateUser)
		authorized.DELETE("/user/:id", controllers.DeleteUser)
	}

	router.Run(":8080")
}
