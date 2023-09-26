package routes

import (
	"add/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(router *gin.Engine, controller controllers.AddController) {
	router.Static("/static", "./static")
	router.GET("/", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "index.html", nil)
    })
	router.POST("/add", controller.Add)
}
