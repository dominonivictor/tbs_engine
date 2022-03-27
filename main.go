package main

import "fmt"

// ENUMS
type DMG_TYPES int8

const (
	DIRECT_DMG_TYPE DMG_TYPES = iota
	TRUE_DMG_TYPE
	NONE_DMG_TYPE
)

type SKILL_TYPES int8

const (
	ATK_SKILL_TYPE SKILL_TYPES = iota
	BUFF_SKILL_TYPE
)

// STRUCTS

type Skill struct {
	name        string
	value       int8
	dmgType     DMG_TYPES
	isFlatValue bool
	skillType   SKILL_TYPES
}

type Stats struct {
	hp    int8
	maxHP int8
	atk   int8
	def   int8
}

type Actor struct {
	name  string
	stats Stats
}

// FUNCTIONS
func validateDmg(dmg int8) int8 {
	if dmg < 0 {
		return int8(0)
	}
	return dmg
}

func dealDmg(target *Actor, dmg int8) {
	if target.stats.hp-dmg < 0 {
		target.stats.hp = 0
	}
	target.stats.hp -= dmg
}

func chooseAtkValue(s Skill, owner *Actor) int8 {
	if s.isFlatValue {
		return s.value
	}
	return owner.stats.atk + s.value

}

func actAtkDirect(s Skill, owner, target *Actor) {
	atk := chooseAtkValue(s, owner)
	def := target.stats.def
	dmg := atk - def
	dmg = validateDmg(dmg)
	dealDmg(target, dmg)
}

func actAtkTrue(s Skill, owner, target *Actor) {
	dmg := chooseAtkValue(s, owner)
	dealDmg(target, dmg)
}

func actAtk(s Skill, owner, target *Actor) {
	var fun func(Skill, *Actor, *Actor)
	switch s.dmgType {
	case DIRECT_DMG_TYPE:
		fun = actAtkDirect
	case TRUE_DMG_TYPE:
		fun = actAtkTrue
	}

	fun(s, owner, target)
}

func chooseSkillFunc(s Skill) func(Skill, *Actor, *Actor) {
	var fun func(Skill, *Actor, *Actor)
	switch s.skillType {
	case ATK_SKILL_TYPE:
		fun = actAtk
	}
	return fun
}

func act(s Skill, owner, target *Actor) {
	fun := chooseSkillFunc(s)
	fun(s, owner, target)
}

func main() {
	fmt.Println("Initializing...")
}
