package utils

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
    comp "github.com/dominonivictor/tbs_engine/components"
)


func Read_csv(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func Load_data() *comp.LoadedData {
  return &comp.LoadedData{
    Reactions_map: Load_reactions(),
    Materials_map: Load_materials(),
    Teams_map: Load_teams(),
    Skills_map: Load_skills(),
  }
}

func Load_materials() map[comp.MAT_ID]comp.Material{
  materials_csv := Read_csv("./data/materials.csv")
  materials := map[comp.MAT_ID]comp.Material{}
  for _, mat := range materials_csv {
    id := comp.MAT_ID(mat[0])
   materials[id] = comp.NewMaterialFromCSV(mat...)
  }
  return materials
}

func Load_reactions() comp.REACTION_MAP {
  reactions_csv := Read_csv("./data/reactions.csv")
  reactions_map := map[comp.MAT_ID]map[comp.I9N_ID]map[comp.MAT_ID]comp.ReactionInfo{
    comp.EMPTY_MATERIAL: map[comp.I9N_ID]map[comp.MAT_ID]comp.ReactionInfo{
      comp.GENERIC_I9N: map[comp.MAT_ID]comp.ReactionInfo{
        comp.EMPTY_MATERIAL: comp.ReactionInfo{
          Value: 1,
          Product: comp.EMPTY_MATERIAL,
        },
      },
    },
  }
  for line_minus_1, info := range reactions_csv {
    interactor_id := comp.MAT_ID(info[0])
    interaction_id := comp.I9N_ID(info[1])
    interacted_id := comp.MAT_ID(info[2])
    value, err := strconv.Atoi(info[3])
    if err != nil {
      fmt.Errorf("Cant convert 4th column (%s) of ./data/reactions.csv Line #%s to int, using default 1", info[3], line_minus_1+1)
      value = 1
    }
    description := info[4]
    product := comp.MAT_ID(info[5])
    reactions_map[interactor_id] = map[comp.I9N_ID]map[comp.MAT_ID]comp.ReactionInfo{
      interaction_id: map[comp.MAT_ID]comp.ReactionInfo{
        interacted_id: comp.NewReactionInfo(value, description, product),
      },
    
    }
  }
  return reactions_map
}

func Load_skills() map[comp.SKILL_ID]comp.Skill {
  skills_csv := Read_csv("./data/skills.csv")
  skills := map[comp.SKILL_ID]comp.Skill{}
  for _, skill := range skills_csv {
    id := comp.SKILL_ID(skill[0])
    skills[id] = comp.NewSkillFromCSV(id)
  }
  return skills
}


func Load_teams() map[comp.TEMPLATE_ID]comp.TeamTemplate{
  teams_csv := Read_csv("./data/teams.csv")
  teams := map[comp.TEMPLATE_ID]comp.TeamTemplate{}
  for _, team := range teams_csv {
    id := comp.TEMPLATE_ID(team[0])
    teams[id] = comp.NewTeamFromCSV(id)
  }
  return teams
}
