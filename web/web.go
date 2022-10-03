package web

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	g := initGin()
	setupRouter(g)
	g.Run(":8080")
}

func initGin() *gin.Engine {
	return gin.Default()
}

func setupRouter(r *gin.Engine) {
	r.POST("/game", postGameController)
	r.GET("/game/:id", getGameController)
	r.GET("/games", getGamesContoller)
}
