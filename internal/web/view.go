package web

import (
	"github.com/resterle/tfiw_go/internal/core/game"
	"github.com/resterle/tfiw_go/internal/core/session"
)

type GameView struct {
	Id       string `json:"id"`
	Self     string `json:"self"`
	Opponent string `json:"opponent"`
	Status   string `json:"status"`
}

func GetGameView(sid string) (GameView, bool) {
	if s, ok := session.Get(sid); ok {
		return MapSession(s, sid), true
	}
	return GameView{}, false
}

func MapSession(s session.Session, sid string) GameView {
	g, gameFound := gameFromSession(s)
	if !gameFound {
		return GameView{}
	}
	pIndex := 0
	if sid == s.GetPlayer2SID() {
		pIndex = 1
	}
	return GameViewFromGame(g, pIndex, sid)
}

func GameViewFromGame(g game.Game, pIndex int, sid string) GameView {
	me := g.Players[pIndex]
	opponent := g.Players[(pIndex+1)%2]
	return GameView{
		Id:       sid,
		Self:     me,
		Opponent: opponent,
		Status:   g.Status,
	}
}

func gameFromSession(s session.Session) (game.Game, bool) {
	if g, ok := game.Get(s.GetGameId()); ok {
		return g, ok
	}
	return game.Game{}, false
}
