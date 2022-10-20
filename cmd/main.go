package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/resterle/tfiw_go/internal/core/game"
	"github.com/resterle/tfiw_go/internal/core/level"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//_, _ := io.ReadAll(resp.Body)

	//fmt.Println(body)

	g := level.NewCave(0, 0, 0)
	l := level.NewLevel()
	l.Fields[0] = level.NewField(&g, 0, 0, false, false)
	l.Fields[1] = level.NewField(&g, 1, 0, false, false)
	l.Fields[2] = level.NewField(&g, 0, 1, false, false)
	l.Fields[3] = level.NewField(&g, 0, 2, false, false)
	//s.Fields = append(s.Fields, game.NewField(&g, 1, 0, false, false))
	//s.Fields = append(s.Fields, game.NewField(&g, 2, 0, false, false))

	fmt.Println(len(l.Fields))
	fmt.Println(l.Fields)
	s, ok := l.Cross(g, 0, 0)
	fmt.Println("######")
	fmt.Println(ok)

	s, ok = s.Cross(g, 0, 0)
	fmt.Println("######")
	fmt.Println(ok)

	n := s.GetNeigbors(s.Fields[0])
	fmt.Println(n)

	dat, err := os.ReadFile("cmd/field.json")
	x := make([]any, 8)

	json.Unmarshal(dat, &x)
	fmt.Println("######")
	fmt.Println(x)

	for _, f := range x {
		t := f.(map[string]any)
		z := parseCave(t)
		//r := z.(int)
		fmt.Println(z)
	}

	//web.Run()
}

func parseCave(m map[string]any) game.Cave {
	x := int(m["x"].(float64))
	y := int(m["y"].(float64))
	points := m["points"].(float64)
	id := (x * 10) + y

	return game.Cave{Id: id, Position: game.Coordinate{X: x, Y: y}, Points: int(points)}
}
