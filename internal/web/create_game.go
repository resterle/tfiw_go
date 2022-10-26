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
	var request CreateGameRequest
	r.ShouldBindJSON(&request)

	newGame, ok := createGame(request.Player)

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

func createGame(playerName string) (*game.Game, bool) {
	if playerName == "" {
		return game.CreateWithDefaultPlayer()
	} else {
		return game.CreateWithPlayer(playerName)
	}
}
