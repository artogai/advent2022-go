package day16

import "testing"

func TestXxx(t *testing.T) {
	eds, rts := parseValves("valves.txt")
	res := countMaxRate(eds, rts)
	t.Log(res)
}
