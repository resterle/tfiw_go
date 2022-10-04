package game

import (
	"github.com/resterle/tfiw_go/core"
)

type Game struct {
	Id      string    `json:"id"`
	Status  string    `json:"status"`
	Players [2]string `json:"players"`
}

var idGenerator func() string = core.RandomId
var games map[string]Game = make(map[string]Game)

func Init(id_genarator func() string) {
	idGenerator = id_genarator
	games = make(map[string]Game)
}

func CreateWithPlayer(player string) (*Game, bool) {
	ok := true
	game := &Game{Id: idGenerator(), Status: "created", Players: [2]string{player}}

	if _, ok := Get(game.Id); ok {
		ok = false
		game = nil
	}
	return game, ok
}

func CreateWithDefaultPlayer() (*Game, bool) {
	return CreateWithPlayer("Player 1")
}

func (g *Game) JoinWithPlayer(player string) {
	g.Players[1] = player
	g.Status = "ready"
}

func (g *Game) JoinWithDefaultPlayer() {
	g.JoinWithPlayer("Player 2")
}

func (g *Game) GetPlayer1() string {
	return g.Players[0]
}

func (g *Game) GetPlayer2() string {
	return g.Players[1]
}
