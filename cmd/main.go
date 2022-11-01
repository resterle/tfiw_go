package main

import (
	"fmt"
	"net/http"

	"github.com/resterle/tfiw_go/internal/core"
	"github.com/resterle/tfiw_go/internal/core/level"
	"go.lair.cx/monads/options"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//_, _ := io.ReadAll(resp.Body)

	//fmt.Println(body)

	l := level.Parse().(level.Level)

	f := l.Caves[3].Fields[7]
	n := l.GetNeigbors(f)
	fmt.Println("--------")
	printField(f)
	fmt.Println("--------")
	for _, g := range n {
		printField(g)
	}

	foo := func(x int) int {
		return x + 1
	}

	bar := func(x int) options.Option[int] {
		if x > 42 {
			return options.Wrap(x)
		}
		return options.Empty[int]()
	}

	opt := options.Wrap(42)
	opt = options.Map(opt, foo)
	opt = options.FlatMap(opt, bar)

	fmt.Println(opt)

	r, err := core.Multi[int, string](
		core.Wrap(10),
		core.Map2(func(i int) (int, error) { return i + 1, nil }),
		core.Map2(func(i int) (string, error) { return fmt.Sprintf("well done %d", i), nil }),
		core.Apply2(func(s string) error {
			fmt.Println(s)
			return nil
		}),
	).Unwrap()
	fmt.Println(err)
	fmt.Println(r)

}

func printField(f level.Field) {
	fmt.Printf("x: %d y: %d\n", f.Position.X, f.Position.Y)
}
