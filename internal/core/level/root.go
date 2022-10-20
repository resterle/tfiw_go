package level

import (
	"fmt"

	"github.com/resterle/tfiw_go/internal/core/game"
)

type Coordinate struct {
	X int
	Y int
}

type Cave struct {
	Id       int
	Position Coordinate
	Points   int
}

type Field struct {
	Cave     *Cave
	Position Coordinate
	Crossed  bool
	Acorn    bool
	Mushroom bool
	Tunnels  [2]*Field
}

type Level struct {
	Fields     []Field
	Hourglass  bool
	FinalTurns int
}

func NewCave(id, x, y int) Cave {
	return Cave{
		Position: Coordinate{X: x, Y: y},
	}
}

func NewField(g *Cave, x, y int, arcon, mushroom bool) Field {
	return Field{
		Cave:     g,
		Position: Coordinate{X: x, Y: y},
		Crossed:  false,
		Acorn:    arcon,
		Mushroom: mushroom,
		Tunnels:  [2]*Field{nil, nil},
	}
}

func (f Field) AddTunnel(n *Field) Field {
	if f.Tunnels[0] == nil {
		f.Tunnels[0] = n
	} else if f.Tunnels[1] == nil {
		f.Tunnels[1] = n
	}
	return f
}

func NewLevel() Level {
	return Level{
		Fields:     make([]Field, 8),
		Hourglass:  false,
		FinalTurns: 3,
	}
}

func (l Level) Cross(c Cave, x, y int) (Level, bool) {
	if f, ok := l.getField(c, x, y); ok {
		if f.Crossed {
			return l, false
		}
		f.Crossed = true
		l.putField(f)
	}
	return l, true
}

func (l Level) getField(c Cave, x, y int) (Field, bool) {
	for _, f := range l.Fields {
		if groupMatch(f, c) && positionMatch(f, x, y) {
			return f, true
		}
	}
	return Field{}, false
}

func (l Level) putField(pf Field) (Level, bool) {
	for i, f := range l.Fields {
		if groupMatch(f, *pf.Cave) && positionMatch(f, pf.Position.X, pf.Position.Y) {
			l.Fields[i] = pf
			return l, true
		}
	}
	return Level{}, false
}

func groupMatch(f Field, g Cave) bool {
	if f.Cave == nil {
		return false
	}
	return f.Cave.Position.X == g.Position.X && f.Cave.Position.X == g.Position.Y
}

func positionMatch(f Field, x, y int) bool {
	return f.Position.X == x && f.Position.Y == y
}

func (l Level) GetNeigbors(f Field) []Field {
	x := f.Position.X
	y := f.Position.Y
	fmt.Println("NN")
	for xOffset := -1; xOffset < 2; xOffset++ {
		for yOffset := -1; yOffset < 2; yOffset++ {
			if xOffset == 0 && yOffset == 0 {
				continue
			}
			nX := x + xOffset
			nY := y + yOffset
			if nX < 0 || nY < 0 {
				continue
			}
			if nf, ok := l.getField(*f.Cave, nX, nY); ok {
				fmt.Println(nf)
			}
		}
	}
	fmt.Println("NN")
	return make([]Field, 1)
}

func parseCave(m map[string]any) game.Cave {
	x := int(m["x"].(float64))
	y := int(m["y"].(float64))
	points := m["points"].(float64)
	id := (x * 10) + y

	return game.Cave{Id: id, Position: game.Coordinate{X: x, Y: y}, Points: int(points)}
}
