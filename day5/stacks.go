package day5

import (
	"file"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func Execute(filename string) []rune {
	commands := readCommands(filename)
	s := stacks{
		stack[rune]{'J', 'H', 'G', 'M', 'Z', 'N', 'T', 'F'},
		stack[rune]{'V', 'W', 'J'},
		stack[rune]{'G', 'V', 'L', 'J', 'B', 'T', 'H'},
		stack[rune]{'B', 'P', 'J', 'N', 'C', 'D', 'V', 'L'},
		stack[rune]{'F', 'W', 'S', 'M', 'P', 'R', 'G'},
		stack[rune]{'G', 'H', 'C', 'F', 'B', 'N', 'V', 'M'},
		stack[rune]{'D', 'H', 'G', 'M', 'R'},
		stack[rune]{'H', 'N', 'M', 'V', 'Z', 'D'},
		stack[rune]{'G', 'N', 'F', 'H'},
	}
	s = s.executeCommands(commands)
	return s.listTop()
}

type stacks []stack[rune]
type command struct{ from, to, count int }

func (s stacks) listTop() []rune {
	return lo.Map(s, func(st stack[rune], _ int) rune {
		return st.peek()
	})
}

func (s stacks) executeCommands(cmds []command) stacks {
	for _, cmd := range cmds {
		s = s.executeCommand(cmd)
	}
	return s
}

func (s stacks) executeCommand(cmd command) stacks {
	nextFromStack, values := s[cmd.from-1].popMany(cmd.count)
	nextToStack := s[cmd.to-1].pushMany(values)
	s[cmd.from-1] = nextFromStack
	s[cmd.to-1] = nextToStack
	return s
}

func readCommands(filename string) []command {
	lines := file.ReadLines(filename)
	commands := lo.Map(lines, func(s string, _ int) command { return parseCommand(s) })
	return commands
}

func parseCommand(s string) command {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	matches := lo.Map(re.FindAllString(s, -1), func(s string, _ int) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	return command{matches[1], matches[2], matches[0]}
}

// naive stack implementation
type stack[A any] []A

func (s stack[A]) pop() (stack[A], A) {
	return s[:len(s)-1], s.peek()
}

func (s stack[A]) push(v A) stack[A] {
	return append(s, v)
}

func (s stack[A]) peek() A {
	return s[len(s)-1]
}

func (s stack[A]) popMany(n int) (stack[A], []A) {
	var currEl A
	buff := make([]A, 0)

	for i := 0; i < n; i++ {
		s, currEl = s.pop()
		buff = append(buff, currEl)
	}

	return s, buff
}

func (s stack[A]) pushMany(values []A) stack[A] {
	for _, v := range values {
		s = s.push(v)
	}
	return s
}
