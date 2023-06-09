package routers

import (
	"book-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBooks)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
