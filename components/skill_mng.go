package components

import (
  "sort"
)


type SkillsMng struct {
  list []Skill
}

func (s Skill) get_bonus() int {
  return 0
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
  interaction_type I9N_ID
}

func NewSkill() Skill {
  defense := NewDefense()
  attack := NewAttack()
  return Skill{
    name: "a skill name",
    defense: defense,
    attack: attack,
    interaction_type: ATK_I9N,
  }
}

func (mng SkillsMng) get_defense_skill_vs(reaction_map REACTION_MAP_TYPE, s Skill, interaction_type I9N_ID) Skill {
  available_defense_skills := mng.filter_defense()
  skills_by_effectiveness_vs_s := mng.choose_best_effectiveness_vs_s(reaction_map, available_defense_skills, s, interaction_type)
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

func (mng SkillsMng) choose_best_effectiveness_vs_s(reaction_map REACTION_MAP_TYPE, def_skills []Skill, s Skill, interaction_type I9N_ID) []Skill {
  sort.Slice(def_skills, func(i, j int) bool {
    return calculate_def_effectiveness(reaction_map, def_skills[i], s) < calculate_def_effectiveness(reaction_map, def_skills[j], s)  
  })
  return def_skills
}


type Attack struct {
  name string
  material Material
}

func NewAttack() Attack {
  return Attack{
    name: "atk name",
    material: NewMaterial(FIRE),
  }
}

type Defense struct {
  name string
  material Material
}

func NewDefense() Defense {
  return Defense{
    name: "defense name",
    material: NewMaterial(WOOD),
  }
}
