package components

import (
  "fmt"
)



type I9N_ID string

const (
  EMPTY_I9N  I9N_ID = "EMPTY_I9N"
  GENERIC_I9N I9N_ID = "GENERIC_I9N"
  ATK_I9N I9N_ID = "ATK_I9N"
  DEF_I9N I9N_ID = "DEF_I9N"
)

type Combat struct {
  skills SkillsMng
  attr   AttrMng
  health HealthMng
  speed  EntitySpeedMng
}

func NewCombat(data *LoadedData, breed Breed) Combat {
    skills := NewSkillsMng(data, breed)
    attr := NewAttrMng(data, breed)
    health := NewHealthMng(data, breed)
    return Combat{
      skills: skills,
      attr: attr,
      health: health,
    }
}

func (c Combat) choose_next_action(){

}


func calculate_def_effectiveness(reaction_map REACTION_MAP, def_skill, atk_skill *Skill) int {
  //var max_effectiveness int // from -2 to 2?, -2 = heals from it, 0 = imune, 2 = double dmg 
  reaction_eff := reaction(
    reaction_map, 
    def_skill.material, 
    atk_skill.material, 
    atk_skill.i9n_id,
  )
  return reaction_eff.Value
}




func choose_target(owner *Entity, targets ...*Entity) *Entity {
  return targets[0]
}

func choose_skill(owner *Entity) *Skill {
  return owner.combat.skills.list[0]
}

func use_skill(reaction_map REACTION_MAP, s *Skill, owner, target *Entity) ActionResultABC {
  bonuses := plan_final_bonuses(reaction_map, s, owner, target)
  t := fmt.Sprintf("%s used skill %s on %s with %d mod", owner.name, s.name, target.name, bonuses)
  fmt.Println("Showing battle result?")
  fmt.Println(t)
  return ActionResult{text: t}
}

const STATUS_BONUS_MOD int = 1
const ATTR_BONUS_MOD int = 1
const SKILL_BONUS_MOD int = 1
const EFFECTIVENESS_BONUS_MOD int = 1

func get_bonuses_for_actor(reaction_map REACTION_MAP,  s, def_s *Skill, ow, tg *Entity, i9n_id I9N_ID) int {
  status_bonus := ow.combat.health.get_bonus()
  attr_bonus  := ow.combat.attr.get_bonus(s)
  skill_bonus := s.get_bonus()
  effectiveness_bonus := reaction(reaction_map, s.material, def_s.material, i9n_id).Value
  return int( 
    status_bonus * STATUS_BONUS_MOD + 
    attr_bonus * ATTR_BONUS_MOD + 
    skill_bonus * SKILL_BONUS_MOD + 
    effectiveness_bonus * EFFECTIVENESS_BONUS_MOD)
  }

func plan_final_bonuses(reaction_map REACTION_MAP, s *Skill, owner, target *Entity) int {
  // 0 = completely misses/does harm to himself, 100 = godlike attack (needs to check lethality and stuff)

  def_s := target.combat.skills.get_defense_skill_vs(reaction_map, s, s.i9n_id)
  // when calculating bonuses you take the atk and defense bonuses, quality of the skills and rank of
  // skills, reaction bonuses with the skills materials, resulting products from the reaction
  bonuses := get_bonuses_for_actor(reaction_map, s, def_s, owner, target, ATK_I9N)
  return bonuses
}



