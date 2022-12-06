package day2

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PlayTournament(filename string) int {
	lines := readFileLines(filename)
	score := 0
	for _, s := range lines {
		game := parseGame(s)
		finished_game := playGame(game)
		score += finished_game.score()
	}
	return score
}

type hand int

const (
	Rock     hand = iota
	Paper         = iota
	Scissors      = iota
)

type outcome int

const (
	Loss outcome = iota
	Draw         = iota
	Win          = iota
)

type game struct{ opponent, player hand }
type finishedGame struct {
	g game
	o outcome
}

func (g finishedGame) score() int {
	return int(g.o)*3 + int(g.g.player) + 1
}

func playGame(g game) finishedGame {
	o := Loss

	if g.player == Rock && g.opponent == Scissors {
		o = Win
	} else if g.player == Scissors && g.opponent == Rock {
		o = Loss
	} else if g.player < g.opponent {
		o = Loss
	} else if g.player > g.opponent {
		o = Win
	} else {
		o = Draw
	}

	return finishedGame{g, o}
}

func parseGame(s string) game {
	return game{parseOpponentHand(s[0]), parsePlayerHand(s[2])}
}

// func parseFinishedGame(s string) game {
// 	return finishedGame{parseOpponentHand(s[0]), parseOutcome(s[2])}
// }

func parseOpponentHand(b byte) hand {
	return parseHand(b, map[byte]hand{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	})
}

func parsePlayerHand(b byte) hand {
	return parseHand(b, map[byte]hand{
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	})
}

func parseHand(b byte, mapping map[byte]hand) hand {
	if mapping == nil {
		panic(fmt.Errorf("mapping is empty"))
	}
	if v, is_defined := mapping[b]; is_defined {
		return v
	} else {
		panic(fmt.Errorf("mapping doesn't contain key: mapping=%v key=%b", mapping, b))
	}
}

func parseOutcome(b byte) outcome {
	switch b {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic(fmt.Errorf("can't parse outcome from %b", b))
}

func readFileLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	return lines
}
