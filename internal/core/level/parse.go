package level

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Option[T any] struct {
	value T
	err   error
}

func Wrap[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		err:   nil,
	}
}

func (x Option[T]) Unwrap() (T, error) {
	return x.value, x.err
}

func Map[T, U any](x Option[T], f func(T) (U, error)) Option[U] {
	if x.err == nil {
		r, err := f(x.value)
		return Option[U]{
			value: r,
			err:   err,
		}
	}
	return Option[U]{
		err: x.err,
	}
}

func MapS[T, U any](x Option[T], f func(T) U) Option[U] {
	if x.err == nil {
		r := f(x.value)
		return Option[U]{
			value: r,
			err:   nil,
		}
	}
	return Option[U]{
		err: x.err,
	}
}

func Apply[T any](x Option[T], f func(T) error) Option[T] {
	if x.err == nil {
		err := f(x.value)
		return Option[T]{
			value: x.value,
			err:   err,
		}
	}
	return x
}

func Parse() any {
	m := read()
	caves := Map(m, parseCaves)
	level, err := MapS(caves, NewLevel).Unwrap()

	if err == nil {
		return level
	}
	return nil
}

func parseCaves(m map[string]any) ([]Cave, error) {
	result := make([]Cave, 0)
	caveArray, ok := m["caves"]
	if !ok {
		return nil, errors.New("could not find key \"cave\"")
	}
	for _, c := range caveArray.([]any) {
		r := parseCave(c.(map[string]any))
		result = append(result, r)
	}
	return result, nil
}

func read() Option[map[string]any] {
	file := Wrap("cmd/field.json")
	data := Map(file, os.ReadFile)
	return Map(data, unmarshal)

}

func unmarshal(data []byte) (map[string]any, error) {
	result := make(map[string]any)
	err := json.Unmarshal(data, &result)
	return result, err
}

func parseCave(m map[string]any) Cave {
	x := int(m["x"].(float64))
	y := int(m["y"].(float64))
	points := m["points"].(float64)
	id := (x * 10) + y

	cave := Cave{Id: id, Position: Coordinate{X: x, Y: y}, Points: int(points)}
	fmt.Println(m)
	fieldArray := m["fields"].([]any)

	fields := make([]Field, 0)
	for _, f := range fieldArray {
		field := parseField(&cave, f.(map[string]any))
		fields = append(fields, field)
	}
	cave.Fields = fields
	return cave
}

func parseField(cave *Cave, m map[string]any) Field {
	x := int(m["x"].(float64))
	y := int(m["y"].(float64))

	return Field{
		Cave:     cave,
		Position: Coordinate{X: x, Y: y},
		Acorn:    false,
		Mushroom: false,
		Crossed:  false,
	}
}
