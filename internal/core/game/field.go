package game

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
	Group    *Cave
	Position Coordinate
	Crossed  bool
	Acorn    bool
	Mushroom bool
	Tunnels  [2]*Field
}

type GameSheet struct {
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
		Group:    g,
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

func CreateGameSheet() GameSheet {
	return GameSheet{
		Fields:     make([]Field, 8),
		Hourglass:  false,
		FinalTurns: 3,
	}
}

func (gs GameSheet) Cross(c Cave, x, y int) (GameSheet, bool) {
	for i, f := range gs.Fields {
		if groupMatch(f, c) && positionMatch(f, x, y) && f.Crossed == false {
			f.Crossed = true
			gs.Fields[i] = f
			return gs, true
		}
	}
	return gs, false
}

func groupMatch(f Field, g Cave) bool {
	if f.Group == nil {
		return false
	}
	return f.Group.Position.X == g.Position.X && f.Group.Position.X == g.Position.Y
}

func positionMatch(f Field, x, y int) bool {
	return f.Position.X == x && f.Position.Y == y
}
