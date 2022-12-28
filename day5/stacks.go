package day5

import (
	"advent2022/file"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

type Stacks []Stack[rune]
type Command struct{ from, to, count int }

func ReadCommands(filename string) []Command {
	return file.ParseFile(filename, parseCommand)
}

func (s Stacks) ListTop() []rune {
	tops := make([]rune, 0, len(s))
	for _, st := range s {
		tops = append(tops, st.Peek())
	}
	return tops
}

func (s Stacks) Execute(cmds []Command, isNewCrane bool) {
	for _, cmd := range cmds {
		s.execute(cmd, isNewCrane)
	}
}

func (s Stacks) execute(cmd Command, isNewCrane bool) {
	values := s[cmd.from].PopMany(cmd.count)
	if isNewCrane {
		values = lo.Reverse(values)
	}
	s[cmd.to].PushMany(values)
}

var cmdRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

func parseCommand(s string) Command {
	matches := lo.Map(cmdRegexp.FindAllString(s, -1), func(s string, _ int) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if i < 0 {
			panic("parseCommand: negative number")
		}
		return i
	})

	return Command{matches[1] - 1, matches[2] - 1, matches[0]}
}
