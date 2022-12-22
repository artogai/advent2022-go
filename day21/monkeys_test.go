package day21

import (
	"fmt"
	"testing"
)

func IgnoreTestExprs(t *testing.T) {
	exprs := NewExprs("monkeys.txt")
	fmt.Println(exprs.Eval("root"))
}

func IgnoreTestExprs2(t *testing.T) {
	exprs := NewExprs("monkeys2.txt")
	evalExprs := exprs.Eval("root")
	fmt.Println(evalExprs.ToString())
	// ((171314231773429322/1617 - x 231040/14553) = 52716091087786)
	// x = 3352886133831
}
