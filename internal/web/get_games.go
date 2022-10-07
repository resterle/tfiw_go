package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resterle/tfiw_go/internal/core/game"
)

func GetGames(r *gin.Context) {
	games := game.All()
	r.JSON(http.StatusOK, games)
}
