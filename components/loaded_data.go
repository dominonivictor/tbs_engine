package components

// the idea is to load these constants from csv files and then
// pass them from main.go into BattleMng and so forth...

type LoadedData struct {
  reaction_map REACTION_MAP
  material_map MATERIAL_MAP
  skill_map SKILL_MAP
  team_map TEMPLATE_MAP
}

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
  t, ok := data.team_map[id]
  if !ok {
    t, _ = data.team_map[DEFAULT_TEAM_ID]
  }
  return t
}
