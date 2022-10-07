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
	r.POST("/game", CreateGame)
	r.GET("/game/:sid", GetGame)
	r.POST("/game/:sid", ChangeGame)
	r.GET("/games", GetGames)
}
