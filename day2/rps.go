package day2

import (
	"advent2022/file"
	"fmt"
)

func PlayTournament(filename string) int {
	lines := file.ReadLines(filename)
	score := 0
	for _, s := range lines {
		game := parseGame(s)
		finishedGame := playGame(game)
		score += finishedGame.score()
	}
	return score
}

func PlayPlannedTournament(filename string) int {
	lines := file.ReadLines(filename)
	score := 0
	for _, s := range lines {
		plannedGame := parsePlannedGame(s)
		playerHand := planHand(plannedGame)
		finishedGame := finishedGame{game{plannedGame.opponent, playerHand}, plannedGame.outcome}
		score += finishedGame.score()
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

type plannedGame struct {
	opponent hand
	outcome  outcome
}

type finishedGame struct {
	game    game
	outcome outcome
}

func (g finishedGame) score() int {
	return int(g.outcome)*3 + int(g.game.player) + 1
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

func planHand(g plannedGame) hand {
	if g.outcome == Loss && g.opponent == Rock {
		return Scissors
	} else if g.outcome == Win && g.opponent == Scissors {
		return Rock
	} else if g.outcome == Loss {
		return g.opponent - 1
	} else if g.outcome == Win {
		return g.opponent + 1
	} else {
		return g.opponent
	}
}

func parseGame(s string) game {
	return game{parseOpponentHand(s[0]), parsePlayerHand(s[2])}
}

func parsePlannedGame(s string) plannedGame {
	return plannedGame{parseOpponentHand(s[0]), parseOutcome(s[2])}
}

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
