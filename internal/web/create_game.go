package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resterle/tfiw_go/internal/core/game"
	"github.com/resterle/tfiw_go/internal/core/session"
)

type CreateGameRequest struct {
	Player string `json:"player"`
}

func CreateGame(r *gin.Context) {
	var newGame *game.Game
	ok := false

	var request CreateGameRequest
	r.BindJSON(&request)

	if playerName := request.Player; playerName == "" {
		newGame, ok = game.CreateWithDefaultPlayer()
	} else {
		newGame, ok = game.CreateWithPlayer(playerName)
	}

	if ok {
		game.Put(*newGame)
		s := session.Create(newGame)
		sid := s.GetPlayer1SID()
		view := MapSession(s, sid)
		r.JSON(http.StatusOK, view)
	} else {
		r.Status(http.StatusInternalServerError)
	}
}
