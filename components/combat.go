package components

import (
  "fmt"
  "sort"
)

type InteractionType string

const (
  ATK_INTERACTION InteractionType = "ATK_INTERACTION"
  DEF_INTERACTION InteractionType = "DEF_INTERACTION"
)

type Combat struct {
  skills SkillsMng
  attr   AttrMng
  health HealthMng
}

func NewCombat() Combat {
    skills := NewSkillsMng()
    attr := NewAttrMng()
    health := NewHealthMng()
    return Combat{
      skills: skills,
      attr: attr,
      health: health,
    }
}

func (c Combat) get_defense_skill_vs(s Skill, interaction_type InteractionType) Skill{
  return c.skills.get_defense_skill_vs(s, interaction_type)
}

type SkillsMng struct {
  list []Skill
}

func (s Skill)get_bonus() float64 {
  return 10
}

func NewSkillsMng() SkillsMng {
  return SkillsMng{
    list: []Skill{
      NewSkill(),
    }, 
  }
}

type Skill struct {
  name string
  defense Defense
  attack Attack
  interaction_type InteractionType
}

func NewSkill() Skill {
  defense := NewDefense()
  attack := NewAttack()
  return Skill{
    name: "a skill name",
    defense: defense,
    attack: attack,
    interaction_type: ATK_INTERACTION,
  }
}

func (mng SkillsMng) get_defense_skill_vs(s Skill, interaction_type InteractionType) Skill {
  available_defense_skills := mng.filter_defense()
  skills_by_effectiveness_vs_s := mng.choose_best_effectiveness_vs_s(available_defense_skills, s, interaction_type)
  return skills_by_effectiveness_vs_s[0] 
}
func (mng SkillsMng) filter_defense() []Skill {
  var result []Skill
  for _, skill := range mng.list {
    if skill.defense.name != "" {
      result = append(result, skill)
    }
  }
  return result
}

func (mng SkillsMng) choose_best_effectiveness_vs_s(def_skills []Skill, s Skill, interaction_type InteractionType) []Skill {
  sort.Slice(def_skills, func(i, j int) bool {
    return calculate_def_effectiveness(def_skills[i], s) < calculate_def_effectiveness(def_skills[j], s)  
  })
  return def_skills
}

func calculate_def_effectiveness(def_skill, atk_skill Skill) float64 {
  //var max_effectiveness float64 // from -2 to 2?, -2 = heals from it, 0 = imune, 2 = double dmg 
  reaction_eff := reaction(def_skill.defense.material, atk_skill.defense.material, atk_skill.interaction_type)
  return reaction_eff
}


type Attack struct {
  name string
  material Material
}

func NewAttack() Attack {
  return Attack{
    name: "atk name",
    material: NewMaterial(),
  }
}

type Defense struct {
  name string
  material Material
}

func NewDefense() Defense {
  return Defense{
    name: "defense name",
    material: NewMaterial(),
  }
}

type AttrMng struct {
  list []Attr

}

func NewAttrMng() AttrMng {
  return AttrMng{
    list: []Attr{
      NewAttr(),
    },
  }
}

func (a AttrMng) get_bonus(s Skill) float64 {
  return 20
}

type Attr struct {
  name string
}

func NewAttr() Attr {
  return Attr{
    name: "attr name",
  }
}

func choose_target(owner Entity, targets ...Entity) Entity {
  return targets[0]
}

func choose_skill(owner Entity) Skill {
  return owner.combat.skills.list[0]
}

func use_skill(s Skill, owner Entity, target Entity) ActionResultABC {
  bonuses := plan_final_bonuses(s, owner, target)
  return ActionResult{text: fmt.Sprintf("%s used skill %s on %s with %q mod", owner.name, s.name, target.name, bonuses)}
}

const STATUS_BONUS_MOD float64 = 15
const ATTR_BONUS_MOD float64 = 15
const SKILL_BONUS_MOD float64 = 15
const EFFECTIVENESS_BONUS_MOD float64 = 15

func get_bonuses_for_actor(s Skill, def_s Skill, ow, tg Entity, interaction_type InteractionType) int {
  status_bonus := ow.combat.health.get_bonus()
  attr_bonus  := ow.combat.attr.get_bonus(s)
  skill_bonus := s.get_bonus()
  effectiveness_bonus := reaction(s.attack.material, def_s.defense.material, interaction_type)
  return int( 
    status_bonus * STATUS_BONUS_MOD + 
    attr_bonus * ATTR_BONUS_MOD + 
    skill_bonus * SKILL_BONUS_MOD + 
    effectiveness_bonus * EFFECTIVENESS_BONUS_MOD)
  }

func plan_final_bonuses(s Skill, owner, target Entity) int {
  // 0 = completely misses/does harm to himself, 100 = godlike attack (needs to check lethality and stuff)

  def_s := target.combat.get_defense_skill_vs(s, s.interaction_type)
  owner_bonuses := get_bonuses_for_actor(s, def_s, owner, target, DEF_INTERACTION)
  target_bonuses := get_bonuses_for_actor(s, def_s, target, owner, DEF_INTERACTION)
  return owner_bonuses - target_bonuses
}



