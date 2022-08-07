package components

import (
  "sort"
)

type SkillsMng struct {
  list []*Skill
}

func (s Skill) get_bonus() int {
  return 0
}

func NewSkillsMng(data *LoadedData, breed Breed) SkillsMng {
  skills := new_skills_from_breed(breed)
  return SkillsMng{
    list: skills,
  }
}

func new_skills_from_breed(breed Breed) []*Skill {
  skills := []*Skill{}
  for _, s := range breed.possible_skills {
    skills = append(skills, NewSkill(NewSkillArgs{id: s.id})) 
  }
  return skills
}

type Skill struct {
  id SKILL_ID
  name string
  defense *Defense
  attack *Attack
  i9n_id I9N_ID
}

type SKILL_MAP map[SKILL_ID]Skill

type NewSkillArgs struct {
  id SKILL_ID
  name string
  defense NewDefenseArgs
  attack NewAttackArgs
  i9n_id I9N_ID
}

func NewSkill(args NewSkillArgs) *Skill {
  defense := NewDefense(args.defense)
  attack := NewAttack(args.attack)
  i9n_id := args.i9n_id
  if i9n_id == EMPTY_I9N {
    i9n_id = GENERIC_I9N
  } 
  return &Skill{
    name: args.name,
    defense: defense,
    attack: attack,
    i9n_id: i9n_id,
  }
}

func (mng SkillsMng) get_defense_skill_vs(reaction_map REACTION_MAP, s *Skill, i9n_id I9N_ID) *Skill {
  available_defense_skills := mng.filter_defense()
  skills_by_effectiveness_vs_s := mng.choose_best_effectiveness_vs_s(reaction_map, available_defense_skills, s, i9n_id)
  return skills_by_effectiveness_vs_s[0] 
}
func (mng SkillsMng) filter_defense() []*Skill {
  var result []*Skill
  for _, skill := range mng.list {
    if skill.defense.name != "" {
      result = append(result, skill)
    }
  }
  return result
}

func (mng SkillsMng) choose_best_effectiveness_vs_s(reaction_map REACTION_MAP, def_skills []*Skill, s *Skill, i9n_id I9N_ID) []*Skill {
  sort.Slice(def_skills, func(i, j int) bool {
    return calculate_def_effectiveness(reaction_map, def_skills[i], s) < calculate_def_effectiveness(reaction_map, def_skills[j], s)  
  })
  return def_skills
}


type Attack struct {
  name string
  material Material
}

type NewAttackArgs struct {
  name string
  mat_id MAT_ID
}

func NewAttack(args NewAttackArgs) *Attack {
  name := args.name
  if name == "" {
    name = "default atk name"
  }
  mat := args.mat_id
  return &Attack{
    name: name,
    material: NewMaterial(mat),
  }
}

type Defense struct {
  name string
  material Material
}

type NewDefenseArgs struct {
  name string
  material Material
}


func NewDefense(args NewDefenseArgs) *Defense {
  name := args.name
  if name == "" {
    name = "default def name"
  }
  mat := args.material
  return &Defense{
    name: name,
    material: NewMaterial(mat.id),
  }
}

type SKILL_ID string
const (
  FIRE_BREATH SKILL_ID = "FIRE_BREATH"
  DEFEND_ALLY SKILL_ID = "DEFEND_ALLY"
  MAKE_IT_RAIN SKILL_ID = "MAKE_IT_RAIN"
  EARTH_WALL SKILL_ID = "EARTH_WALL"
  VACCUM_SUCK SKILL_ID = "VACCUM_SUCK"
)

type ACTION_ID string
const (
  ATK_ACT_ID ACTION_ID = "ATK_ACT_ID"
  DEF_ACT_ID ACTION_ID = "DEF_ACT_ID"
  POISON_ACT_ID ACTION_ID = "POISON_ACT_ID"
  BURN_ACT_ID ACTION_ID = "BURN_ACT_ID"
)

func NewSkillFromCSV(id SKILL_ID) Skill {
  return Skill{}
}
