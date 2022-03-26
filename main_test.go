package example

import (
	"testing"
)

func TestSimpleAtk(t *testing.T) {
	stat := Stats{hp: 10, hpMax: 10, atk: 1, def: 1}
	a1 := Actor{name: "a1", stats: stat}
	a2 := Actor{name: "a2", stats: stat}
	s := Skill{name: "s1", dmgValue: 5, dmgType: DIRECT}

	act(s, &a1, &a2)

	if a2.stats.hp != 5 {
		t.Errorf("Actor hp hit by skill should be 5, was %v", a2)
	}
}
