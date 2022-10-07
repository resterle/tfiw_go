package game

import (
	"testing"
)

const expectedDefaultPlayer1 = "Player 1"
const expectedDefaultPlayer2 = "Player 2"

var stubbed_id string = "id"
var idGeneratorStub func() string = func() string {
	return stubbed_id
}

func TestCreateWithDefaultPlayer(t *testing.T) {
	Init(idGeneratorStub)

	stubbed_id = "id"
	if actual, ok := CreateWithDefaultPlayer(); testTrue(t, ok) {
		assertId(t, *actual, stubbed_id)
		assertPlayer1(t, *actual, expectedDefaultPlayer1)
		assertPlayer2(t, actual, "")
	}
}

func TestCreateWithPlayer(t *testing.T) {
	Init(idGeneratorStub)

	expectedPlayer1 := "Hans"

	stubbed_id = "id"
	if actual, ok := CreateWithPlayer(expectedPlayer1); testTrue(t, ok) {
		assertId(t, *actual, stubbed_id)
		assertPlayer1(t, *actual, expectedPlayer1)
		assertPlayer2(t, actual, "")
	}
}

func TestPut(t *testing.T) {
	Init(idGeneratorStub)

	stubbed_id = "id"
	game, _ := CreateWithDefaultPlayer()
	if ok := Put(*game); testTrue(t, ok) {
		assertNumberOfGames(t, 1)
	}

}

func TestPutMultiGames(t *testing.T) {
	Init(idGeneratorStub)

	stubbed_id = "id_1"
	game, _ := CreateWithDefaultPlayer()
	if ok := Put(*game); testTrue(t, ok) {
		assertNumberOfGames(t, 1)
	}
	stubbed_id = "id_2"
	game, _ = CreateWithDefaultPlayer()
	if ok := Put(*game); testTrue(t, ok) {
		assertNumberOfGames(t, 2)
	}

}

func TestPutWithoutId(t *testing.T) {
	Init(idGeneratorStub)

	stubbed_id = ""
	game, _ := CreateWithDefaultPlayer()
	if ok := Put(*game); ok {
		t.Error("Expected Put to fail")
	}
}

func TestUpdate(t *testing.T) {
	Init(idGeneratorStub)

	givenStatus := "updated"

	stubbed_id = "id"
	game, _ := CreateWithDefaultPlayer()
	Put(*game)

	// Put should save a copy and therefore not change the supplied game
	game.Status = givenStatus
	actual, _ := Get(game.Id)
	if actualStatus := actual.Status; actualStatus != "created" {
		t.Errorf("expected status to be \"created\" but it was %s", actualStatus)
	}

	// A put with the same game should update it
	if ok := Put(*game); testTrue(t, ok) {
		assertNumberOfGames(t, 1)
		actual, _ := Get(game.Id)
		if actualStatus := actual.Status; actualStatus != "updated" {
			t.Errorf("expected status to be \"updated\" but it was %s", actualStatus)
		}
	}

}

func TestGet(t *testing.T) {
	Init(idGeneratorStub)

	stubbed_id = "id"
	game, _ := CreateWithDefaultPlayer()
	Put(*game)

	if actual, ok := Get(stubbed_id); testTrue(t, ok) {
		assertId(t, actual, stubbed_id)
	}
}

func TestGetNotFound(t *testing.T) {
	Init(idGeneratorStub)
	_, ok := Get(stubbed_id)
	testFalse(t, ok)
}

func TestDelete(t *testing.T) {
	Init(idGeneratorStub)
	game, _ := CreateWithDefaultPlayer()
	Put(*game)
	if ok := Delete(*game); testTrue(t, ok) {
		assertNumberOfGames(t, 0)
	}
}

func TestDeleteNonExistent(t *testing.T) {
	Init(idGeneratorStub)
	game, _ := CreateWithDefaultPlayer()
	if ok := Delete(*game); ok {
		t.Error("Expected Delte to fail but it succeeded")
	}
}

func assertNumberOfGames(t *testing.T, expected int) {
	if actual := len(games); actual != expected {
		t.Errorf("assert number of games: got %d, wanted %d", actual, expected)
	}
}

func assertId(t *testing.T, game Game, expected string) {
	if actual := game.Id; actual != expected {
		t.Errorf("assert id: got %s, wanted %s", actual, expected)
	}
}

func assertPlayer1(t *testing.T, game Game, expected string) {
	if actual := game.Players[0]; actual != expected {
		t.Errorf("assert player 1: got %s, wanted %s", actual, expected)
	}
}

func assertPlayer2(t *testing.T, game *Game, expected string) {
	if actual := game.Players[1]; actual != expected {
		t.Errorf("assert player 2: got %s, wanted %s", actual, expected)
	}
}

func testTrue(t *testing.T, b bool) bool {
	if !b {
		t.Error("got false wanted true")
	}
	return b
}

func testFalse(t *testing.T, b bool) bool {
	if b {
		t.Error("got true wanted false")
	}
	return b
}
