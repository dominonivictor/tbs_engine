package main

import (
	//tbs "tbs_engine/main"
	"testing"
)

func TestAtkSkills(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := &Actor{name: "a1", stats: stat}
	target := &Actor{name: "a2", stats: stat}
	t.Run("Direct Atk without Flat Dmg", func(t *testing.T) {
		d := Dmg{value: 4, dmgType: DIRECT_DMG_TYPE}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 3
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("Direct Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 4, dmgType: DIRECT_DMG_TYPE, isFlatValue: true}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 2
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("True Atk without Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 5, dmgType: TRUE_DMG_TYPE}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 6
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
	t.Run("True Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 5, dmgType: TRUE_DMG_TYPE, isFlatValue: true}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 5
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
}

func TestStatusStatBuff(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := &Actor{name: "a1", stats: stat}
	buff := Buff{
		name:         "b1",
		buffType:     STAT_BUFF_TYPE,
		timer:        2,
		value:        -2,
		statAffected: ATK_STAT,
	}
	s := Skill{
		name:      "s1",
		skillType: BUFF_SKILL_TYPE,
		buff:      buff,
	}
	t.Run("Status Effect, check target has status", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}

		act(s, owner, target)

		wantStatus := "b1"
		gotStatus := target.statuses[buff.name].name
		if gotStatus != wantStatus {
			t.Errorf("Actor status is not apropriate, got %v, want %v", gotStatus, wantStatus)
		}
	})
	t.Run("Status Effect, check target has stat (de)buffed", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}

		act(s, owner, target)

		gotActorAtk := target.stats.atk
		wantActorAtk := -1
		if gotActorAtk != wantActorAtk {
			t.Errorf("Actor atk stat is not right, got %v, want %v", gotActorAtk, wantActorAtk)
		}
	})
	t.Run("Status Effect, check target has only 1 status, even if hit twice", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}

		act(s, owner, target)
		act(s, owner, target)

		gotStatusesLen := len(target.statuses)
		wantStatusesLen := 1
		if gotStatusesLen != wantStatusesLen {
			t.Errorf("Actor statuses len is not right, got %v, want %v", gotStatusesLen, wantStatusesLen)
		}
		//passTurn(battle)

	})
}

func TestStatusPassiveBuff(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := &Actor{name: "a1", stats: stat}
	atkSkill := Skill{
		name:      "s1",
		skillType: BUFF_SKILL_TYPE,
		dmg:       Dmg{value: 10, dmgType: DIRECT_DMG_TYPE},
	}
	t.Run("Status Effect, passive perfect defense", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}
		perfectDefenseSkill := Skill{
			name: "perfect defense",
			buff: Buff{name: "perfect defense", value: 0, timer: 1, buffType: PASSIVE_BUFF_TYPE, passiveId: "perfect_defense"},
		}

		act(perfectDefenseSkill, owner, target)
		act(atkSkill, owner, target)

		wantHP := 10
		gotHP := target.stats.hp
		if gotHP != wantHP {
			t.Errorf("Actor wrong HP, got %v, want %v", gotHP, wantHP)
		}
	})
}
