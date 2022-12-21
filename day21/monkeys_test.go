package day21

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	exprs := NewExprs("monkeys.txt")
	fmt.Println(exprs.Eval("root"))
}
