package components

import (
  "fmt"
)


type BattleMng struct{
  team1 Team
  team2 Team
  terrain Terrain
  reports string
}

type Team struct {
  actors []Entity
  name string
}

func (t Team) are_all_actors_inactive() bool {
  at_least_one_is_active := false
  for _, actor := range t.actors {
    if actor.is_active(){
      at_least_one_is_active = true
    }
  }
  return at_least_one_is_active
}

func (t Team) choose_actor() Entity {
  return t.actors[0]
}

type Terrain struct {

}


func choose_and_act(t1, t2 Team){
    attacker := t1.choose_actor()
    skill := choose_skill(attacker)
    target := choose_target(attacker, t2.actors[0]) 
    result := use_skill(skill, attacker, target)
    fmt.Println(result.String())
}

func (b BattleMng) start_battle(){
  // THIS IS CRAP FOR NOW
  for !b.team1.are_all_actors_inactive() || !b.team2.are_all_actors_inactive(){
    // here the ideal scenario would be to take the speed of actors and choose next to move
    choose_and_act(b.team1, b.team2)
    choose_and_act(b.team2, b.team1)

  }

}
