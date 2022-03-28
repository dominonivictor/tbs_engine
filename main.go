package main

import (
	"fmt"
	"math"
)

// ENUMS
type DMG_TYPES int

const (
	DIRECT_DMG_TYPE DMG_TYPES = iota
	TRUE_DMG_TYPE
	NONE_DMG_TYPE
)

type SKILL_TYPES int

const (
	ATK_SKILL_TYPE SKILL_TYPES = iota
	BUFF_SKILL_TYPE
)

type BUFF_TYPES int

const (
	STAT_BUFF_TYPE BUFF_TYPES = iota
	DOT_OR_HOT_BUFF_TYPE
	PASSIVE_BUFF_TYPE
)

type STATS int

const (
	ATK_STAT STATS = iota
	DEF_STAT
)

// STRUCTS

type Buff struct {
	name         string
	value        int
	timer        int
	buffType     BUFF_TYPES
	statAffected STATS
	passiveId    string
}

type Dmg struct {
	value       int
	dmgType     DMG_TYPES
	isFlatValue bool
}

type Skill struct {
	name      string
	skillType SKILL_TYPES
	dmg       Dmg
	buff      Buff
}

type Stats struct {
	hp    int
	maxHP int
	atk   int
	def   int
}

type Actor struct {
	name     string
	stats    Stats
	statuses map[string]*Buff
}

// FUNCTIONS
func validateDmg(dmg int) int {
	if dmg < 0 {
		return int(0)
	}
	return dmg
}

func dealDmg(target *Actor, dmg int) {
	if target.stats.hp-dmg < 0 {
		target.stats.hp = 0
	}
	target.stats.hp -= dmg
}

func calculateAtkValue(s Skill, owner *Actor) int {
	if s.dmg.isFlatValue {
		return s.dmg.value
	}
	return owner.stats.atk + s.dmg.value

}

func actAtkDirect(s Skill, owner, target *Actor) {
	atk := calculateAtkValue(s, owner)
	def := target.stats.def
	dmg := atk - def
	dmg = validateDmg(dmg)
	dealDmg(target, dmg)
}

func actAtkTrue(s Skill, owner, target *Actor) {
	dmg := calculateAtkValue(s, owner)
	dealDmg(target, dmg)
}

func actAtk(s Skill, owner, target *Actor) {
	var fun func(Skill, *Actor, *Actor)
	switch s.dmg.dmgType {
	case DIRECT_DMG_TYPE:
		fun = actAtkDirect
	case TRUE_DMG_TYPE:
		fun = actAtkTrue
	}

	fun(s, owner, target)
}

func addStatus(target *Actor, status *Buff) {
	currStatus, found := target.statuses[status.name]

	if !found {
		target.statuses[status.name] = status
	} else {
		currStatus.timer = int(math.Max(float64(currStatus.timer), float64(status.timer)))
	}
}

func applyBuffStat(b Buff, target *Actor) {
	switch b.statAffected {
	case ATK_STAT:
		target.stats.atk += b.value
	case DEF_STAT:
		target.stats.def += b.value
	}
}

func actBuffStat(b Buff, owner, target *Actor) {
	status := &b
	addStatus(target, status)
	applyBuffStat(b, target)
}

func actBuff(s Skill, owner, target *Actor) {
	var fun func(Buff, *Actor, *Actor)
	b := s.buff
	switch b.buffType {
	case STAT_BUFF_TYPE:
		fun = actBuffStat
	}
	fun(b, owner, target)
}

func act(s Skill, owner, target *Actor) {
	// for each non-empty component, execute proper function
	if s.dmg != (Dmg{}) {
		actAtk(s, owner, target)
	}
	if s.buff != (Buff{}) {
		actBuff(s, owner, target)
	}
}

func main() {
	fmt.Println("Initializing...")
	a := &Actor{statuses: make(map[string]*Buff)}
	s := &Buff{name: "status1", timer: 1}
	addStatus(a, s)
	fmt.Printf("Actor: %+v, status: %+v", a, s)

}
