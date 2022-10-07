package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resterle/tfiw_go/internal/core/game"
	"github.com/resterle/tfiw_go/internal/core/session"
)

func ChangeGame(r *gin.Context) {
	var p GameView
	r.BindJSON(&p)

	sid := r.Params.ByName("sid")

	if s, ok := session.Get(sid); ok {
		httpStatus := http.StatusOK
		pIndex := 0
		if sid == s.GetPlayer2SID() {
			pIndex = 1
		}
		if og, ok := game.Get(s.GetGameId()); ok {
			g := og
			if !(handleSelf(pIndex, &p, &og, &g, &httpStatus) && handleStatus(&p, &og, &g, &httpStatus)) {
				g = og
			}
			game.Put(g)
			r.JSON(httpStatus, GameViewFromGame(g, pIndex, sid))
			return
		}
	}
	r.Status(http.StatusNotFound)

}

func handleSelf(pIndex int, p *GameView, og *game.Game, g *game.Game, httpStatus *int) bool {
	if p.Self == "" {
		return true
	}
	if og.Status != "created" {
		*httpStatus = http.StatusPreconditionFailed
		return false
	}
	setPlayerFunc := og.SetPlayer1
	if pIndex == 1 {
		setPlayerFunc = og.SetPlayer2
	}
	*g = setPlayerFunc(p.Self)
	*httpStatus = http.StatusOK
	return true
}

func handleStatus(p *GameView, og *game.Game, g *game.Game, httpStatus *int) bool {
	if p.Status == og.Status {
		return true
	}
	if p.Status == "ready" {
		ug, err := g.Start()
		if err != nil {
			*httpStatus = http.StatusPreconditionFailed
			return false
		}
		*g = ug
	}
	return true
}
