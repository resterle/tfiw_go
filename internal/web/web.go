package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(initGin),
	fx.Invoke(setupRouter),
	fx.Invoke(startup),
)

func startup(g *gin.Engine) {
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
