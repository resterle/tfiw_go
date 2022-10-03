package game

func Get(id string) (*Game, bool) {
	res, ok := games[id]
	return &res, ok
}

func Put(game Game) bool {
	res := false
	if game.Id != "" {
		games[game.Id] = game
		res = true
	}
	return res
}

func Delete(game Game) bool {
	res := false
	if _, ok := Get(game.Id); ok {
		delete(games, game.Id)
		res = true
	}
	return res
}

func All() []Game {
	r := make([]Game, 0, len(games))
	for _, v := range games {
		r = append(r, v)
	}
	return r
}
