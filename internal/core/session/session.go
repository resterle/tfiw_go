package session

import (
	"github.com/resterle/tfiw_go/internal/core"
	"github.com/resterle/tfiw_go/internal/core/game"
)

type Session struct {
	gameId     string
	player1SID string
	player2SID string
}

func Create(game *game.Game) Session {
	s := Session{
		gameId:     game.Id,
		player1SID: core.RandomId(),
		player2SID: core.RandomId(),
	}
	Put(s)
	return s
}

func (s *Session) GetPlayer1SID() string {
	return s.player1SID
}

func (s *Session) GetPlayer2SID() string {
	return s.player2SID
}

func (s *Session) GetGameId() string {
	return s.gameId
}
