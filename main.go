package main

import (
	"github.com/resterle/tfiw_go/core/game"
	"github.com/resterle/tfiw_go/web"
)

func main() {
	g, _ := game.CreateWithPlayer("Hans")
	game.Put(*g)

	web.Run()
}
