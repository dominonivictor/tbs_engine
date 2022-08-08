package components

import (
  "sort"
  "strconv"
  "fmt"
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
  rank RANKS
  material *Material
  defense *Defense
  attack *Attack
  i9n_id I9N_ID
}


type NewSkillArgs struct {
  id SKILL_ID
  name string
  rank RANKS
  material NewMaterialArgs
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

type RANKS int
// 8 High master
// 7 master
// 6 expert
// 5 pro
// 4 advanced
// 3 average
// 2 familiar
// 1 novice
// 0 dabbling
const (
  LEGENDARY_5 RANKS = 10
  LEGENDARY RANKS = 9
  HIGH_MASTER RANKS = 8
  MASTER RANKS = 7
  EXPERT RANKS = 6
  PRO RANKS = 5
  ADVANCED RANKS = 4
  AVERAGE RANKS = 3
  FAMILIAR RANKS = 2
  NOVICE RANKS = 1
  DABBLING RANKS = 0
)

type Attack struct {
  name string
  value int
}

// ranks
// 10 legendary 5 max dmg 20? 
// 9 legendary max dmg 18?
// 8 High master
// 7 master
// 6 expert
// 5 professional
// 4 advanced
// 3 average
// 2 familiar
// 1 novice
// 0 dabbling

type NewAttackArgs struct {
  name string
  value int
}

func NewAttack(args NewAttackArgs) *Attack {
  name := args.name
  if name == "" {
    name = "default atk name"
  }
  value := args.value
  if value == 0 {
    value = 1
  }
  return &Attack{
    name: name,
    value: value,
  }
}

func NewAttackFromCSV(id string, value int) NewAttackArgs {
  return NewAttackArgs{
    name: id,
    value: value,
  }
}

type Defense struct {
  name string
  value int
}

type NewDefenseArgs struct {
  name string
  value int
}


func NewDefense(args NewDefenseArgs) *Defense {
  name := args.name
  if name == "" {
    name = "default def name"
  }
  return &Defense{
    name: name,
    value: args.value,
  }
}

func NewDefenseFromCSV(id string, value int) NewDefenseArgs {
  return NewDefenseArgs{
    name: id,
    value: value,
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

func NewSkillFromCSV(headers map[string]int, skill_row []string, material_map MATERIAL_MAP) NewSkillArgs {
  rank, err := strconv.Atoi(skill_row[headers["rank"]])
  if err != nil {
    fmt.Errorf("Cant convert #%d column (%s) of ./data/skills.csv to int, using default 0", headers["rank"], skill_row[headers["rank"]])
    rank = 0
  }
  atk_value, err2 := strconv.Atoi(skill_row[headers["atk_value"]])
  if err2 != nil {
    fmt.Errorf("Cant convert #%d column (%s) of ./data/skills.csv to int, using default 0", headers["rank"], skill_row[headers["rank"]])
    atk_value = 0
  }
  def_value, err3 := strconv.Atoi(skill_row[headers["def_value"]])
  if err3 != nil {
    fmt.Errorf("Cant convert #%d column (%s) of ./data/skills.csv to int, using default 0", headers["rank"], skill_row[headers["rank"]])
    def_value = 0
  }
  fmt.Printf("header: %+v, Skill_row[]: %+v\n", headers, skill_row)
  mat_id := MAT_ID(skill_row[headers["mat_id"]])
  fmt.Printf("MATERIAL ID: %s\n", string(mat_id))
  mat_args := material_map[mat_id]
  fmt.Printf("MATERIAL ARGS: %+v\n", mat_args)
  return NewSkillArgs {
    id: SKILL_ID(skill_row[headers["id"]]),
    name: skill_row[headers["name"]],
    rank: RANKS(rank),
    material: mat_args,
    i9n_id: I9N_ID(skill_row[headers["i9n_id"]]),
    attack: NewAttackFromCSV(skill_row[headers["atk_id"]], atk_value),
    defense: NewDefenseFromCSV(skill_row[headers["def_id"]], def_value),
  }
}
