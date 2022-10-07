package session

import "fmt"

var playerLookup map[string]Session = make(map[string]Session)

func Get(sid string) (Session, bool) {

	fmt.Println(sid)
	fmt.Println("---")
	for k, _ := range playerLookup {
		fmt.Println(k)
	}

	if res, ok := playerLookup[sid]; ok {
		return res, ok
	}
	return Session{}, false
}

func Put(s Session) {
	playerLookup[s.player1SID] = s
	playerLookup[s.player2SID] = s
}
