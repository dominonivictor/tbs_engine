package main

import (
	//tbs "tbs_engine/main"
	"testing"
)

func TestAtkSkills(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := Actor{name: "a1", stats: stat}
	target := Actor{name: "a2", stats: stat}
	t.Run("Direct Atk without Flat Dmg", func(t *testing.T) {
		s := Skill{name: "s1", value: 4, dmgType: DIRECT_DMG_TYPE}

		act(s, &owner, &target)

		want := target.stats.maxHP - 3
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("Direct Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		s := Skill{name: "s1", value: 4, dmgType: DIRECT_DMG_TYPE, isFlatValue: true}

		act(s, &owner, &target)

		want := target.stats.maxHP - 2
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("True Atk without Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		s := Skill{name: "s1", value: 5, dmgType: TRUE_DMG_TYPE}

		act(s, &owner, &target)

		want := target.stats.maxHP - 6
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
	t.Run("True Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		s := Skill{name: "s1", value: 5, dmgType: TRUE_DMG_TYPE, isFlatValue: true}

		act(s, &owner, &target)

		want := target.stats.maxHP - 5
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
}
