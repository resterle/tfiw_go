package level

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/resterle/tfiw_go/internal/core"
)

func Parse() any {
	m := read()
	caves := core.Map(m, parseCaves)
	level, err := core.MapS(caves, NewLevel).Unwrap()

	if err == nil {
		return level
	}
	fmt.Printf("===> error: %s <===\n", err.Error())
	return nil
}

func parseCaves(m map[string]any) ([]Cave, error) {
	result := make([]Cave, 0)
	caveArray, ok := m["caves"]
	if !ok {
		return nil, errors.New("could not find key \"cave\"")
	}
	for _, c := range caveArray.([]any) {
		r, err := parseCave(c.(map[string]any))
		if err != nil {
			return result, err
		}
		result = append(result, r)
	}
	return result, nil
}

func read() core.Option[map[string]any] {
	return core.Multi[string, map[string]any](
		core.Wrap("cmd/field.json"),
		core.Map2(os.ReadFile),
		core.Map2(unmarshal),
	)
}

func unmarshal(data []byte) (map[string]any, error) {
	result := make(map[string]any)
	err := json.Unmarshal(data, &result)
	return result, err
}

func parseCave(m map[string]any) (Cave, error) {

	wm := core.Wrap(m)
	wm = core.Map(wm, as_int_func("x"))
	wm = core.Map(wm, as_int_func("y"))
	wm = core.Map(wm, as_int_func("points"))

	wc := core.MapS(wm, func(m map[string]any) Cave {
		x := get_int(m, "x")
		y := get_int(m, "y")
		points := get_int(m, "points")
		id := (x * 10) + y

		return Cave{Id: id, Position: Coordinate{X: x, Y: y}, Points: int(points), Fields: make([]Field, 0)}
	})

	fa := m["fields"].([]any)
	for _, f := range fa {
		wc = core.Map(wc, func(cave Cave) (Cave, error) {
			field, err := parseField(&cave, f.(map[string]any))
			cave.Fields = append(cave.Fields, field)
			return cave, err
		})
	}

	return wc.Unwrap()
}

func maybe_as_int(m map[string]any, k string) (map[string]any, error) {
	v, ok := m[k]
	if ok {
		m[k] = int(v.(float64))
		return m, nil
	}
	return m, fmt.Errorf("key %s not found", k)
}

func as_int_func(key string) func(map[string]any) (map[string]any, error) {
	return func(m map[string]any) (map[string]any, error) {
		return maybe_as_int(m, key)
	}
}

func get_int(m map[string]any, k string) int {
	return m[k].(int)
}

func parseField(cave *Cave, m map[string]any) (Field, error) {

	wm := core.Wrap(m)
	wm = core.Map(wm, as_int_func("x"))
	wm = core.Map(wm, as_int_func("y"))

	m, err := wm.Unwrap()

	if err == nil {
		field := Field{
			Cave:     cave,
			Position: Coordinate{X: m["x"].(int), Y: m["y"].(int)},
			Acorn:    false,
			Mushroom: false,
			Crossed:  false,
		}
		return field, nil
	}

	return Field{}, err

}
