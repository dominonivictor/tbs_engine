package components

import (
  "fmt"
)


type BattleMng struct{
  Team1 Team
  Team2 Team
  terrain Terrain
  reports string
}

type Team struct {
  actors []Entity
  name string
}

func NewTeam() Team {
  actors := []Entity{NewEntity()}
  return Team{
    actors: actors,
    name: "team name",
  }
}

func (t Team) is_at_least_one_actor_active() bool {
  at_least_one_is_active := false
  for _, actor := range t.actors {
    if actor.is_active(){
      at_least_one_is_active = true
    }
  }
  fmt.Printf("at_least_one_is_active: %t\n", at_least_one_is_active)
  //return at_least_one_is_active
  return true
}

func (t Team) choose_actor() Entity {
  return t.actors[0]
}

type Terrain struct {

}

func choose_and_act(t1, t2 Team){
    fmt.Println("choosing and acting")
    attacker := t1.choose_actor()
    skill := choose_skill(attacker)
    target := choose_target(attacker, t2.actors[0]) 
    result := use_skill(skill, attacker, target)
    fmt.Println(result.String())
    fmt.Println("finished choosing and acting")
}

func (b BattleMng) Start_battle(){
  fmt.Println("Starting battle...")
  // THIS IS CRAP FOR NOW
  c1 := b.Team1.is_at_least_one_actor_active()  
  c2 := b.Team2.is_at_least_one_actor_active()
  for c1 && c2 {
    fmt.Println("inside loop")
    // here the ideal scenario would be to take the speed of actors and choose next to move
    choose_and_act(b.Team1, b.Team2)
    choose_and_act(b.Team2, b.Team1)
    //c1 = !b.Team1.is_at_least_one_actor_active()  
    c1 = false
    c2 = !b.Team2.is_at_least_one_actor_active()  
  }
  fmt.Println("Battle is over!")
}
