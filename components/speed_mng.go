package components


type BattleSpeedMng struct {
  current_turn uint
  battle *BattleMng
}

func NewSpeedMng() BattleSpeedMng {
  return BattleSpeedMng{
    current_turn: uint(0),
  }
}

func (s BattleSpeedMng) pass_time(battle BattleMng) {

  for _, actor := range battle.entities {
    actor.combat.speed.actor_pass_time()
    if actor.combat.speed.is_actor_ready_to_act(){
      actor.combat.choose_next_action()
    }
  }
}

type EntitySpeedMng struct {
  time_to_act int 
}


func (s EntitySpeedMng) actor_pass_time(){

}

func (s EntitySpeedMng) is_actor_ready_to_act() bool {
  return true
}
