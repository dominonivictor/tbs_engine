package main

import (
	"testing"
)

func TestAtkSkills(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2, mag: 10, res: 5}
	owner := &Actor{name: "a1", stats: stat, statuses: make(map[string]*Buff)}
	target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}
	t.Run("Direct Physical Atk without Flat Dmg", func(t *testing.T) {
		d := Dmg{value: 4, dmgType: DIRECT_DMG_TYPE, dmgCategory: PHYSICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 3
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("Direct Physical Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 4, dmgType: DIRECT_DMG_TYPE, isFlatValue: true, dmgCategory: PHYSICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 2
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("True Atk without Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 5, dmgType: TRUE_DMG_TYPE, dmgCategory: PHYSICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 6
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
	t.Run("True Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 5, dmgType: TRUE_DMG_TYPE, isFlatValue: true, dmgCategory: PHYSICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 5
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}

	})
	t.Run("Direct Magical Atk without Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 4, dmgType: DIRECT_DMG_TYPE, dmgCategory: MAGICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 9
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("Direct Magical Atk with Flat Dmg", func(t *testing.T) {
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 6, dmgType: DIRECT_DMG_TYPE, isFlatValue: true, dmgCategory: MAGICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := target.stats.maxHP - 1
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
	t.Run("Take dmg while having a Shield (temporary?)", func(t *testing.T) {
		target.stats.shieldHP = 5
		target.stats.hp = target.stats.maxHP
		d := Dmg{value: 10, dmgType: DIRECT_DMG_TYPE, isFlatValue: true, dmgCategory: PHYSICAL_DMG}
		s := Skill{name: "s1", dmg: d}

		act(s, owner, target)

		want := 7
		if target.stats.hp != want {
			t.Errorf("Actor hp hit by skill should be %v, was %+v", want, target)
		}
	})
}

func TestStatusStatBuff(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := &Actor{name: "a1", stats: stat, statuses: make(map[string]*Buff)}
	buff := Buff{
		name:         "b1",
		buffType:     STAT_BUFF_TYPE,
		timer:        2,
		value:        -2,
		statAffected: ATK_STAT,
		buffId:       "minus_2_atk",
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
		gotStatus := target.statuses[buff.buffId].name
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
	owner := &Actor{name: "a1", stats: stat, statuses: make(map[string]*Buff)}
	atkSkill := Skill{
		name:      "s1",
		skillType: BUFF_SKILL_TYPE,
		dmg:       Dmg{value: 10, dmgType: DIRECT_DMG_TYPE, dmgCategory: PHYSICAL_DMG},
	}
	t.Run("Status Effect, passive perfect defense", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}
		perfectDefenseSkill := Skill{
			name: "perfect defense",
			buff: Buff{name: "perfect defense", value: 0, timer: 1, buffType: PASSIVE_BUFF_TYPE, buffId: "perfect_defense"},
		}

		act(perfectDefenseSkill, target, target)
		act(atkSkill, owner, target)

		wantHP := 10
		gotHP := target.stats.hp
		if gotHP != wantHP {
			t.Errorf("Actor wrong HP, got %v, want %v", gotHP, wantHP)
		}
	})
	t.Run("Status Effect, freeze makes unable to use skill", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff)}
		freezeSkill := Skill{
			name: "freeze",
			buff: Buff{name: "freeze", value: 0, timer: 1, buffType: PASSIVE_BUFF_TYPE, buffId: "freeze"},
		}

		act(freezeSkill, target, owner)
		act(atkSkill, owner, target)

		wantHP := 10
		gotHP := target.stats.hp
		if gotHP != wantHP {
			t.Errorf("Actor wrong HP, got %v, want %v", gotHP, wantHP)
		}
	})
}

func TestElementalEffectiveness(t *testing.T) {
	stat := Stats{hp: 10, maxHP: 10, atk: 1, def: 2}
	owner := &Actor{name: "a1", stats: stat}
	atkSkill := Skill{
		name:      "s1",
		skillType: BUFF_SKILL_TYPE,
		dmg:       Dmg{value: 2, dmgType: DIRECT_DMG_TYPE, dmgCategory: PHYSICAL_DMG},
		element:   WATER,
	}
	t.Run("Water Element Skill is 150% effective vs Fire target, vs DIRECT_DMG", func(t *testing.T) {
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff), element: FIRE}

		act(atkSkill, owner, target)

		wantHP := 6
		gotHP := target.stats.hp
		if gotHP != wantHP {
			t.Errorf("Actor wrong HP, got %v, want %v", gotHP, wantHP)
		}
	})
	t.Run("Water Element Skill is 150% effective vs Fire target, vs TRUE_DMG", func(t *testing.T) {
		atkSkill.dmg.dmgType = TRUE_DMG_TYPE
		target := &Actor{name: "a2", stats: stat, statuses: make(map[string]*Buff), element: FIRE}

		act(atkSkill, owner, target)

		wantHP := 4
		gotHP := target.stats.hp
		if gotHP != wantHP {
			t.Errorf("Actor wrong HP, got %v, want %v", gotHP, wantHP)
		}
	})
}

func TestAutoBattleManager(t *testing.T) {
	t.Run("Stronger team wins vs Weaker team", func(t *testing.T) {
		t1 := &Team{name: "t1", powerValue: 20}
		t2 := &Team{name: "t2", powerValue: 10}

		results := autoBattle(*t1, *t2)
		wantWinner := t1.name
		gotWinner := results.winner.name

		if gotWinner != wantWinner {
			t.Errorf("Wrong winning team, got: %+v, want: %+v", gotWinner, wantWinner)
		}
	})
}
