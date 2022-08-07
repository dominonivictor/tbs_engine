package components

import (
  "fmt"
)


type BattleMng struct{
  Team1 *Team
  Team2 *Team
  speed *BattleSpeedMng
  terrain *TerrainMng
  reports string
  entities []*Entity
}

type NewBattleMngArgs struct {
  Team1 *Team
  Team2 *Team
  speed *BattleSpeedMng
  terrain *TerrainMng
  reports string
  entities []*Entity
}

func NewBattleMng(data *LoadedData, args NewBattleMngArgs) *BattleMng {
  t1 := args.Team1
  if t1 == nil {
    t1 = NewTeam(data,  data.Teams_map[DEFAULT_TEAM_ID])
  }
  t2 := args.Team2
  if t2 == nil {
    t2 = NewTeam(data,  data.Teams_map[DEFAULT_TEAM_ID])
  }

  return &BattleMng{
    Team1: t1,
    Team2: t2,
    entities: args.entities,
    terrain: args.terrain,
    speed: args.speed,
  }
}


type Team struct {
  actors []*Entity
  name string
}


func NewTeam(data *LoadedData, template TeamTemplate) *Team {
  // create default team from template
  actors := []*Entity{NewEntity(data, NewBreed(template.breeds[0]))}
  return &Team{
    actors: actors,
    name: "team name",
  }
}

func NewTeamFromCSV(tid TEMPLATE_ID) TeamTemplate {
  return TeamTemplate{}
}

func (t Team) is_at_least_one_actor_active() bool {
  at_least_one_is_active := false
  for _, actor := range t.actors {
    if actor.is_active(){
      at_least_one_is_active = true
    }
  }
  fmt.Printf("at_least_one_is_active: %t\n", at_least_one_is_active)
  return at_least_one_is_active
}

func (t Team) choose_actor() *Entity {
  return t.actors[0]
}

type TerrainMng struct {

}

func choose_and_act(reaction_map REACTION_MAP, t1, t2 *Team){
    fmt.Println("choosing and acting")
    attacker := t1.choose_actor()
    skill := choose_skill(attacker)
    target := choose_target(attacker, t2.actors[0]) 
    result := use_skill(reaction_map, skill, attacker, target)
    fmt.Println(result.String())
    fmt.Println("finished choosing and acting")
}

func (b BattleMng) Start_battle_till_death_or_other_interruption(reaction_map REACTION_MAP){
  fmt.Println("Starting battle...")
  // THIS IS CRAP FOR NOW
  c1 := b.Team1.is_at_least_one_actor_active()  
  c2 := b.Team2.is_at_least_one_actor_active()
  for c1 && c2 {
    fmt.Println("inside loop")
    // here the ideal scenario would be to take the speed of actors and choose next to move
    b.speed.pass_time(b)
    _ = b.next_actor_act(reaction_map, b.Team1, b.Team2)
    //fmt.Println(action_result.print())
    //choose_and_act(reaction_map, b.Team1, b.Team2)
    //choose_and_act(reaction_map, b.Team2, b.Team1)
    c1 = b.Team1.is_at_least_one_actor_active()
    c2 = b.Team2.is_at_least_one_actor_active()  
  }
  fmt.Println("Battle is over!")
}

func (b BattleMng) next_actor_act(reaction_map REACTION_MAP, t1, t2 *Team) ActionResult {
  choose_and_act(reaction_map, b.Team1, b.Team2)
  return ActionResult{}
  
}

