package components

// the idea is to load these constants from csv files and then
// pass them from main.go into BattleMng and so forth...

type LoadedData struct {
  Reactions_map REACTION_MAP
  Materials_map MATERIAL_MAP
  Skills_map SKILL_MAP
  Teams_map TEMPLATE_MAP
}

type SKILL_MAP map[SKILL_ID]NewSkillArgs
type REACTION_MAP map[MAT_ID]map[I9N_ID]map[MAT_ID]ReactionInfo 
type MATERIAL_MAP map[MAT_ID]NewMaterialArgs

type TEMPLATE_MAP map[TEMPLATE_ID]TeamTemplate
type TEMPLATE_ID string
const (
  DEFAULT_TEAM_ID TEMPLATE_ID = "DEFAULT_TEAM_ID"
  LATEGAME_3_DWARV_PARTY TEMPLATE_ID = "LATEGAME_3_DWARV_PARTY"
  LATEGAME_DRAGON_LICH_DWARF_NECRO TEMPLATE_ID = "LATEGAME_DRAGON_LICH_DWARF_NECRO"
)

type TeamTemplate struct {
  breeds []BREED_ID
}

func NewTeamTemplate(data LoadedData, id TEMPLATE_ID) TeamTemplate {
  t, ok := data.Teams_map[id]
  if !ok {
    t, _ = data.Teams_map[DEFAULT_TEAM_ID]
  }
  return t
}
