package file

import (
	"log"
	"os"
	"strings"
)

func ParseFile[A any](filename string, parse func(s string) A) []A {
	lines := ReadFileLines(filename)
	parsed := make([]A, 0, len(lines))
	for _, line := range lines {
		parsed = append(parsed, parse(line))
	}
	return parsed
}

func ReadFileLines(filename string) []string {
	content := ReadFile(filename)

	lines := strings.Split(string(content), "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
