package ops

import (
	"strings"
)

// Board describes the current state of the game
type Board struct {
	xLen int
	yLen int
	pCt  int
	ps   []Planet
	ss   [][]Ship
}

// MakeBoard from a slice of game state tokens
func MakeBoard(xLen, yLen int, gameData string) Board {
	tokens := strings.Split(gameData, " ")
	pCt := readTokenInt(tokens, 0)
	tokens = tokens[1:]

	b := Board{
		xLen: xLen,
		yLen: yLen,
		ss:   make([][]Ship, pCt),
	}

	for k := range b.ss {
		curID := readTokenInt(tokens, 0)
		if curID != k {
			panic("curID is not iteration when making board")
		}

		shipCt := readTokenInt(tokens, 1)
		tokens = tokens[2:]

		for i := 0; i < shipCt; i++ {
			s, trimmedTokens := MakeShip(k, tokens)
			tokens = trimmedTokens

			b.ss[k] = append(b.ss[k], s)
		}
	}

	plntCt := readTokenInt(tokens, 0)
	tokens = tokens[1:]

	for i := 0; i < plntCt; i++ {
		p, trimmedTokens := MakePlanet(tokens)
		tokens = trimmedTokens

		b.ps = append(b.ps, p)
	}

	return b
}

// Dimensions ...
func (b *Board) Dimensions() (int, int) {
	return b.xLen, b.yLen
}

// PlayerCt ...
func (b *Board) PlayerCt() int {
	return b.pCt
}

// Planets ...
func (b *Board) Planets() []Planet {
	return b.ps
}

// Ships ...
func (b *Board) Ships() [][]Ship {
	return b.ss
}
