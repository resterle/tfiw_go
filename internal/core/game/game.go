package game

import (
	"errors"

	"github.com/resterle/tfiw_go/internal/core"
)

type Game struct {
	Id         string       `json:"id"`
	Status     string       `json:"status"`
	Players    [2]string    `json:"players"`
	GameSheets [2]GameSheet `json:"game_sheets"`
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

func (g Game) JoinWithPlayer(player string) Game {
	g.Players[1] = player
	g.Status = "ready"
	return g
}

func (g Game) JoinWithDefaultPlayer() Game {
	return g.JoinWithPlayer("Player 2")
}

func (g Game) GetPlayer1() string {
	return g.Players[0]
}

func (g Game) GetPlayer2() string {
	return g.Players[1]
}

func (g Game) SetPlayer1(p string) Game {
	g.Players[0] = p
	return g
}

func (g Game) SetPlayer2(p string) Game {
	g.Players[1] = p
	return g
}

func (g Game) Start() (Game, error) {
	if g.GetPlayer1() == "" || g.GetPlayer2() == "" {
		err := errors.New("players not joined")
		return g, err
	}
	if g.Status != "created" {
		err := errors.New("transition not allowed")
		return g, err
	}
	g.Status = "ready"
	return g, nil
}
