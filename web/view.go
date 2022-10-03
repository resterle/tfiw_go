package web

import (
	"fmt"

	"github.com/resterle/tfiw_go/core"
	"github.com/resterle/tfiw_go/core/game"
)

var mappings map[string]*GameViewMapping = make(map[string]*GameViewMapping)

type GameView struct {
	Id       string `json:"id"`
	Self     string `json:"self"`
	Opponent string `json:"opponent"`
	Status   string `json:"status"`
}

type GameViewMapping struct {
	GameId    string
	Player1Id string
	Player2Id string
}

func CreateMapping(game *game.Game) *GameViewMapping {
	mapping := newMapping(game)
	putMapping(*mapping)
	return mapping
}

func GetGameView(id *string) (*GameView, bool) {
	fmt.Println(mappings)
	if mapping, ok := getMapping(id); ok {
		g, _ := game.Get(mapping.GameId)
		me := g.Players[0]
		opponent := g.Players[1]
		if id == &mapping.Player2Id {
			me = g.Players[1]
			opponent = g.Players[0]
		}
		return &GameView{
			Id:       *id,
			Self:     me,
			Opponent: opponent,
			Status:   g.Status,
		}, true
	}
	return nil, false
}

func getMapping(id *string) (*GameViewMapping, bool) {
	if res, ok := mappings[*id]; ok {
		return res, ok
	}
	return nil, false
}

func newMapping(game *game.Game) *GameViewMapping {
	return &GameViewMapping{
		GameId:    game.Id,
		Player1Id: core.RandomId(),
		Player2Id: core.RandomId(),
	}
}

func putMapping(mapping GameViewMapping) {
	mappings[mapping.Player1Id] = &mapping
	mappings[mapping.Player2Id] = &mapping
}

func deleteMapping(mapping *GameViewMapping) {

}
