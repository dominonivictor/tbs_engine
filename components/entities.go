package components

type Entity struct {
  name string
  sprite string
  combat Combat
}

func (e Entity) is_active() bool {
  return e.combat.health.is_active()
}

func NewEntity(data *LoadedData, breed Breed) *Entity {
  combat := NewCombat(data, breed)
  return &Entity{
    name: "entity name",
    sprite: "E",
    combat: combat,
  }
}

type Breed struct {
  race BREED_ID
  possible_skills []NewSkillArgs
}

type NewBreedArgs struct {
  race BREED_ID
  possible_skills []NewSkillArgs
}

func NewBreed(args NewBreedArgs) Breed {
  return Breed{
    race: args.race,
    possible_skills: args.possible_skills,
  }
}

func NewBreedFromCSV(args ...string) NewBreedArgs {
  
}

type BREED_MAP map[BREED_ID]Breed 

type BREED_ID string
const (
  DRAGONS BREED_ID = "Dragons"
  ELVES BREED_ID = "Elves"
  DWARVES BREED_ID = "Dwarves"
  GNOMES BREED_ID = "Gnomes"
  GOBLINS BREED_ID = "Goblins"
  ORCS BREED_ID = "Orcs"
)


  //t1 := comp.NewTeam(data, data.Template_map[comp.LATEGAME_3_DWARV_PARTY])
  //t2 := comp.NewTeam(data, data.Template_map[comp.LATEGAME_DRAGON_LICH_DWARF_NECRO])
