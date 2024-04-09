package internal

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	h := Handlers{}

	r.GET("/", h.Index)

	g := r.Group("/books")
	{
		g.GET("/:ID", h.GetBookByIDHandler)
		g.POST("/", h.AddNewBookHandler)
		g.DELETE("/:ID", h.DeleteBookHandler)
	}
}
