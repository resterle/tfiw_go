package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resterle/tfiw_go/core/game"
)

func getGamesContoller(r *gin.Context) {
	games := game.All()
	r.JSON(http.StatusOK, games)
}

type NewGameParams struct {
	Player string `json:"player"`
}

func postGameController(r *gin.Context) {
	var newGame *game.Game
	ok := false

	var params NewGameParams
	r.BindJSON(&params)

	if playerName := params.Player; playerName != "" {
		newGame, ok = game.CreateWithPlayer(playerName)
	} else {
		newGame, ok = game.CreateWithDefaultPlayer()
	}

	if ok {
		game.Put(*newGame)
		mapping := CreateMapping(newGame)
		view, _ := GetGameView(&mapping.Player1Id)
		r.JSON(http.StatusOK, view)
	} else {
		r.Status(http.StatusInternalServerError)
	}
}

func getGameController(r *gin.Context) {
	id := r.Params.ByName("id")

	if v, ok := GetGameView(&id); ok {
		r.JSON(http.StatusOK, v)
	} else {
		r.Status(http.StatusNotFound)
	}
}
