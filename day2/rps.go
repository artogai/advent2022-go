package day2

import (
	"advent2022/file"
	"fmt"
)

func PlayTournament(filename string) int {
	games := file.ParseFile(filename, parseGame)
	score := 0
	for _, g := range games {
		finishedGame := playGame(g)
		score += finishedGame.score()
	}
	return score
}

func PlayPlannedTournament(filename string) int {
	plans := file.ParseFile(filename, parseGamePlan)
	score := 0
	for _, p := range plans {
		player := playerHand(p.opponent, p.outcome)
		finishedGame := finishedGame{game{p.opponent, player}, p.outcome}
		score += finishedGame.score()
	}
	return score
}

type hand int

const (
	rock     hand = iota
	paper    hand = iota
	scissors hand = iota
)

type outcome int

const (
	loss outcome = iota
	draw outcome = iota
	win  outcome = iota
)

type game struct{ opponent, player hand }

type finishedGame struct {
	game
	outcome outcome
}

type gamePlan struct {
	opponent hand
	outcome  outcome
}

func (g finishedGame) score() int {
	return int(g.outcome)*3 + int(g.game.player) + 1
}

func playGame(g game) finishedGame {
	o := loss

	if g.player == rock && g.opponent == scissors {
		o = win
	} else if g.player == scissors && g.opponent == rock {
		o = loss
	} else if g.player < g.opponent {
		o = loss
	} else if g.player > g.opponent {
		o = win
	} else {
		o = draw
	}

	return finishedGame{g, o}
}

func playerHand(opponent hand, outcome outcome) hand {
	if outcome == loss && opponent == rock {
		return scissors
	} else if outcome == win && opponent == scissors {
		return rock
	} else if outcome == loss {
		return opponent - 1
	} else if outcome == win {
		return opponent + 1
	} else {
		return opponent
	}
}

var (
	opponentHandMapping = map[byte]hand{
		'A': rock,
		'B': paper,
		'C': scissors,
	}
	playerHandMapping = map[byte]hand{
		'X': rock,
		'Y': paper,
		'Z': scissors,
	}
	outcomeMapping = map[byte]outcome{
		'X': loss,
		'Y': draw,
		'Z': win,
	}
)

func parseGame(s string) game {
	return game{parseOpponentHand(s[0]), parsePlayerHand(s[2])}
}

func parseGamePlan(s string) gamePlan {
	return gamePlan{parseOpponentHand(s[0]), parseOutcome(s[2])}
}

func parseOpponentHand(b byte) hand {
	return parseByte(b, opponentHandMapping)
}

func parsePlayerHand(b byte) hand {
	return parseByte(b, playerHandMapping)
}

func parseOutcome(b byte) outcome {
	return parseByte(b, outcomeMapping)
}

func parseByte[A any](b byte, mapping map[byte]A) A {
	if v, exists := mapping[b]; exists {
		return v
	}
	panic(fmt.Errorf("mapping doesn't contain key: mapping=%v key=%b", mapping, b))
}
