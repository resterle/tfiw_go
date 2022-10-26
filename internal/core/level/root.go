package level

type Coordinate struct {
	X int
	Y int
}

type Cave struct {
	Id       int
	Position Coordinate
	Points   int
	Fields   []Field
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
	Caves      []Cave
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

func NewLevel(caves []Cave) Level {
	return Level{
		Caves:      caves,
		Hourglass:  false,
		FinalTurns: 3,
	}
}

func (l Level) Cross(cid int, x, y int) (Level, bool) {
	if f, ok := l.getField(cid, x, y); ok {
		if f.Crossed {
			return l, false
		}
		f.Crossed = true
		l.putField(f)
	}
	return l, true
}

func (l Level) getField(cid int, x, y int) (Field, bool) {
	for _, cave := range l.Caves {
		if cave.Id == cid {
			for _, f := range cave.Fields {
				if positionMatch(f, x, y) {
					return f, true
				}
			}
		}
	}
	return Field{}, false
}

func (l Level) putField(pf Field) (Level, bool) {
	cave := pf.Cave
	for i, f := range cave.Fields {
		if positionMatch(f, pf.Position.X, pf.Position.Y) {
			cave.Fields[i] = pf
			return l, true
		}
	}
	return Level{}, false
}

func positionMatch(f Field, x, y int) bool {
	return f.Position.X == x && f.Position.Y == y
}

func (l Level) GetNeigbors(f Field) []Field {
	result := make([]Field, 0)
	offsets := [2]int{-1, 1}

	for _, offset := range offsets {
		if nf, ok := l.getField(f.Cave.Id, f.Position.X+offset, f.Position.Y); ok {
			result = append(result, nf)
		}
		if nf, ok := l.getField(f.Cave.Id, f.Position.X, f.Position.Y+offset); ok {
			result = append(result, nf)
		}
	}

	return result
}

func (l Level) GetNeigbors2(f Field) []Field {
	result := make([]Field, 0)
	x := f.Position.X
	y := f.Position.Y
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
			if nf, ok := l.getField(f.Cave.Id, nX, nY); ok {
				result = append(result, nf)
			}
		}
	}
	return result
}
