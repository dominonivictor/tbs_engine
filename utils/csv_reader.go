package utils

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
    comp "github.com/dominonivictor/tbs_engine/components"
)


func Read_csv(filePath string) (map[string]int, [][]string) {
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
    headers := make(map[string]int)
    for index, col_name := range records[0] {
      headers[col_name] = index
    }

    return headers, records[1:]
}

func Load_data() *comp.LoadedData {
  //reacts := Load_reactions()
  mats := Load_materials()
  skills := Load_skills(mats)
  return &comp.LoadedData{
    //Reactions_map: Load_reactions(),
    Materials_map: mats,
    //Teams_map: Load_teams(),
    Skills_map: skills,
  }
}

//type MATERIAL_MAP map[MAT_ID]NewMaterialArgs
func Load_materials() comp.MATERIAL_MAP {
  _, materials_csv := Read_csv("./data/materials.csv")
  materials := comp.MATERIAL_MAP(map[comp.MAT_ID]comp.NewMaterialArgs{})
  for _, mat := range materials_csv {
    id := comp.MAT_ID(mat[0])
    materials[id] = comp.NewMaterialFromCSV(mat...)
  }
  return materials
}

func Load_breeds() comp.BREED_MAP {
  _, breeds_csv := Read_csv("./data/breeds.csv")
  breeds := comp.MATERIAL_MAP(map[comp.MAT_ID]comp.NewMaterialArgs{})
  for _, breed := range breeds_csv {
    id := comp.MAT_ID(breed[0])
    breeds[id] = comp.NewBreedFromCSV(breed...)
  }
  return breeds
}

func Load_reactions() comp.REACTION_MAP {
  _, reactions_csv := Read_csv("./data/reactions.csv")
  reactions_map := map[comp.MAT_ID]map[comp.I9N_ID]map[comp.MAT_ID]comp.ReactionInfo{
    comp.VOID: map[comp.I9N_ID]map[comp.MAT_ID]comp.ReactionInfo{
      comp.GENERIC_I9N: map[comp.MAT_ID]comp.ReactionInfo{
        comp.VOID: comp.ReactionInfo{
          Value: 1,
          Product: comp.VOID,
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

func Load_skills(material_map comp.MATERIAL_MAP) map[comp.SKILL_ID]comp.NewSkillArgs {
  fmt.Printf("Loading skills...")
  headers, skills_csv := Read_csv("./data/skills.csv")
  skills := map[comp.SKILL_ID]comp.NewSkillArgs{}
  for _, skill_row := range skills_csv {
    id_idx := headers["id"]
    id := comp.SKILL_ID(skill_row[id_idx])
    skills[id] = comp.NewSkillFromCSV(headers, skill_row, material_map)
  }
  return skills
}


func Load_teams() map[comp.TEMPLATE_ID]comp.TeamTemplate{
  _, teams_csv := Read_csv("./data/teams.csv")
  teams := map[comp.TEMPLATE_ID]comp.TeamTemplate{}
  for _, team := range teams_csv {
    id := comp.TEMPLATE_ID(team[0])
    teams[id] = comp.NewTeamFromCSV(id)
  }
  return teams
}
