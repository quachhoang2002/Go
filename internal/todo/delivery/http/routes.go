package http

import "github.com/gin-gonic/gin"

// MapRoutes maps the routes to the handler functions
func (h handler) MapRoutes(r *gin.RouterGroup) {
	r.DELETE("/:id", h.delete)
	r.GET("", h.list)
	r.POST("", h.add)
}
