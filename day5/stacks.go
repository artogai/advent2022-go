package day5

import (
	"advent2022/file"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

type Stacks []Stack[rune]
type Command struct{ from, to, count int }

func (s Stacks) ListTop() []rune {
	return lo.Map(s, func(st Stack[rune], _ int) rune {
		return st.Peek()
	})
}

func (s Stacks) ExecuteCommandsFromFile(filename string, isNewCrane bool) {
	cmds := readCommands(filename)
	s.ExecuteCommands(cmds, isNewCrane)
}

func (s Stacks) ExecuteCommands(cmds []Command, isNewCrane bool) {
	for _, cmd := range cmds {
		s.ExecuteCommand(cmd, isNewCrane)
	}
}

func (s Stacks) ExecuteCommand(cmd Command, isNewCrane bool) {
	values := s[cmd.from].PopMany(cmd.count)
	if isNewCrane {
		values = lo.Reverse(values)
	}
	s[cmd.to].PushMany(values)
}

func readCommands(filename string) []Command {
	lines := file.ReadLines(filename)
	commands := lo.Map(lines, func(s string, _ int) Command { return parseCommand(s) })
	return commands
}

func parseCommand(s string) Command {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	matches := lo.Map(re.FindAllString(s, -1), func(s string, _ int) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	return Command{matches[1] - 1, matches[2] - 1, matches[0]}
}

// naive stack implementation
type Stack[A any] []A

func (s *Stack[A]) Push(v A) {
	*s = append(*s, v)
}

func (s *Stack[A]) Pop() A {
	ds := *s
	res := ds[len(ds)-1]
	*s = ds[:len(ds)-1]
	return res
}

func (s *Stack[A]) Peek() A {
	res := (*s)[len(*s)-1]
	return res
}

func (s *Stack[A]) PushMany(values []A) {
	for _, v := range values {
		s.Push(v)
	}
}

func (s *Stack[A]) PopMany(n int) []A {
	buff := make([]A, 0, n)
	for i := 0; i < n; i++ {
		buff = append(buff, s.Pop())
	}

	return buff
}
