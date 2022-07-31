package main

import (
	"fmt"
	"math"
)

// CONSTANTS

var elementsTable map[ELEMENTS]map[ELEMENTS]float64 = map[ELEMENTS]map[ELEMENTS]float64{
	FIRE: map[ELEMENTS]float64{
		WATER: 0.5,
	},
	WATER: map[ELEMENTS]float64{
		FIRE: 2,
	}}

// ENUMS
type DMG_TYPES int

const (
	DIRECT_DMG_TYPE DMG_TYPES = iota
	TRUE_DMG_TYPE
	NONE_DMG_TYPE
)

type DMG_CATEGORIES int

const (
	PHYSICAL_DMG DMG_CATEGORIES = iota
	MAGICAL_DMG
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

type STATS_ENUM int

const (
	ATK_STAT STATS_ENUM = iota
	DEF_STAT
)

type ELEMENTS string

const (
	FIRE  ELEMENTS = "FIRE"
	WATER          = "WATER"
)

// STRUCTS

type Buff struct {
	name         string
	value        int
	timer        int
	buffType     BUFF_TYPES
	statAffected STATS_ENUM
	buffId       string
}

type Dmg struct {
	value       int
	dmgType     DMG_TYPES
	isFlatValue bool
	dmgCategory DMG_CATEGORIES
}

type Skill struct {
	name      string
	skillType SKILL_TYPES
	dmg       Dmg
	buff      Buff
	element   ELEMENTS
}

type Stats struct {
	hp       int
	maxHP    int
	shieldHP int
	//maxShieldHP int
	atk int
	def int
	mag int
	res int
}

type Actor struct {
	name     string
	stats    Stats
	statuses map[string]*Buff
	element  ELEMENTS
}

type Team struct {
	name       string
	powerValue int
	actors     []*Actor
}

type AutoBattleResult struct {
	winner Team
	loser  Team
	isDraw bool
}

// METHODS

func (actor *Actor) hasBuff(buffId string) bool {
	_, ok := actor.statuses[buffId]
	return ok
}

func (actor *Actor) canAct() bool {
	return !actor.hasBuff("freeze")
}

// FUNCTIONS
func validateDmg(dmg int, target *Actor) int {
	if dmg <= 0 || target.hasBuff("perfect_defense") {
		return 0
	}
	return dmg
}

func dealDmg(target *Actor, dmg int) {
	if target.stats.shieldHP >= dmg {
		target.stats.shieldHP -= dmg
	} else {
		effectiveDmg := dmg - target.stats.shieldHP
		target.stats.shieldHP = 0
		if target.stats.hp-effectiveDmg < 0 {
			target.stats.hp = 0
		}
		target.stats.hp -= effectiveDmg
	}
}

func addElementalEffectiveness(dmg int, s Skill, target *Actor) int {
	elementMultiplier, found := elementsTable[s.element][target.element]
	if found {
		dmgWithMultiplier := int(elementMultiplier * float64(dmg))
		return dmgWithMultiplier
	}
	return dmg
}

func calculateAtkValue(s Skill, owner, target *Actor) int {
	dmg := s.dmg.value
	if s.dmg.isFlatValue {
		return dmg
	}
	if s.dmg.dmgCategory == PHYSICAL_DMG {

		return owner.stats.atk + dmg
	} else if s.dmg.dmgCategory == MAGICAL_DMG {

		return owner.stats.mag + dmg
	}
	return 0

}

func calculateDefValue(s Skill, owner, target *Actor) int {
	if s.dmg.dmgCategory == PHYSICAL_DMG {

		return owner.stats.def
	} else if s.dmg.dmgCategory == MAGICAL_DMG {

		return owner.stats.res
	}
	return 0
}

func actAtkDirect(s Skill, owner, target *Actor) {
	atk := calculateAtkValue(s, owner, target)
	atkWithBonus := addElementalEffectiveness(atk, s, target)
	def := calculateDefValue(s, owner, target)
	dmg := atkWithBonus - def
	dmg = validateDmg(dmg, target)
	dealDmg(target, dmg)
}

func actAtkTrue(s Skill, owner, target *Actor) {
	dmg := calculateAtkValue(s, owner, target)
	dmgWithBonus := addElementalEffectiveness(dmg, s, target)
	dealDmg(target, dmgWithBonus)
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
	currStatus, found := target.statuses[status.buffId]

	if !found {
		target.statuses[status.buffId] = status
	} else {
		currStatus.timer = int(math.Max(float64(currStatus.timer), float64(status.timer)))
	}
}

func applyBuffStat(b *Buff, target *Actor) {
	switch b.statAffected {
	case ATK_STAT:
		target.stats.atk += b.value
	case DEF_STAT:
		target.stats.def += b.value
	}
}

func actBuffStat(b *Buff, owner, target *Actor) {
	addStatus(target, b)
	applyBuffStat(b, target)
}

func actPassiveStat(b *Buff, owner, target *Actor) {
	addStatus(target, b)
}

func actBuff(s Skill, owner, target *Actor) {
	var fun func(*Buff, *Actor, *Actor)
	b := &s.buff
	switch b.buffType {
	case STAT_BUFF_TYPE:
		fun = actBuffStat
	case PASSIVE_BUFF_TYPE:
		fun = actPassiveStat
	}
	fun(b, owner, target)
}

func act(s Skill, owner, target *Actor) {
	// for each non-empty component, execute proper function
	if owner.canAct() {
		if s.dmg != (Dmg{}) {
			actAtk(s, owner, target)
		}
		if s.buff != (Buff{}) {
			actBuff(s, owner, target)
		}
	}
}

func calculateActorPowerLevel(actor Actor) float32 {
	s := actor.stats
	return 0.75*(float32(s.maxHP)+0.5*float32(s.def)) + float32(s.atk)
}

func calculateTeamPowerLevel(team Team) float32 {
	// TODO: Take into consideration resistances, skills, elemental mods
	var teamPowerLevel float32
	for _, actor := range team.actors {
		teamPowerLevel += calculateActorPowerLevel(*actor)
	}
	return teamPowerLevel
}

func autoBattle(t1, t2 Team) *AutoBattleResult {
	t1Power := calculateTeamPowerLevel(t1)
	t2Power := calculateTeamPowerLevel(t2)
	if t1Power == t2Power {
		return &AutoBattleResult{isDraw: true}
	} else if t1Power > t2Power {

		return &AutoBattleResult{winner: t1, loser: t2, isDraw: false}
	} else {

		return &AutoBattleResult{winner: t2, loser: t1, isDraw: false}
	}
}

func main() {
	fmt.Println("Initializing...")
	a := &Actor{statuses: make(map[string]*Buff)}
	s := &Buff{name: "status1", timer: 1}
	addStatus(a, s)
	fmt.Printf("Actor: %+v, status: %+v", a, s)

}
