package components


var reactions_map map[MAT_NAME]map[InteractionType]map[MAT_NAME]float64 = map[MAT_NAME]map[InteractionType]map[MAT_NAME]float64{
  WOOD: map[InteractionType]map[MAT_NAME]float64{
    ATK_INTERACTION: map[MAT_NAME]float64{
      WOOD: 1,
    },
  },
  FIRE: map[InteractionType]map[MAT_NAME]float64{
    ATK_INTERACTION: map[MAT_NAME]float64{
      WOOD: 2,
    },
  },
}

type MAT_NAME string
const (
  FIRE MAT_NAME = "FIRE"
  WATER MAT_NAME = "WATER"
  EARTH MAT_NAME = "EARTH"
  AIR MAT_NAME = "AIR"
  LIGHTNING MAT_NAME = "LIGHTNING"
  WOOD MAT_NAME = "WOOD"
  CLOTH MAT_NAME = "CLOTH"
  BONE MAT_NAME = "BONE"
  FAT MAT_NAME = "FAT"
  FEATHER MAT_NAME = "FEATHER"
  STONE MAT_NAME = "STONE"
  MUSCLE MAT_NAME = "MUSCLE"
  VAPOR MAT_NAME = "VAPOR"
  RUBBER MAT_NAME = "RUBBER"
  HAIR MAT_NAME = "HAIR"
  METAL MAT_NAME = "METAL"
  SCALE MAT_NAME = "SCALE"
  LEATHER MAT_NAME = "LEATHER"
)

type MAT_STATUS string
const (
  BURN_MAT_STATUS MAT_STATUS = "BURN_MAT_STATUS"
  CHARGED_MAT_STATUS MAT_STATUS = "CHARGED_MAT_STATUS"
  WET_MAT_STATUS MAT_STATUS = "WET_MAT_STATUS"
  DRY_MAT_STATUS MAT_STATUS = "DRY_MAT_STATUS"
  FROZEN_MAT_STATUS MAT_STATUS = "FROZEN_MAT_STATUS"
  COLD_MAT_STATUS MAT_STATUS = "COLD_MAT_STATUS"
  HOT_MAT_STATUS MAT_STATUS = "HOT_MAT_STATUS"
  WEBBED_MAT_STATUS MAT_STATUS = "WEBBED_MAT_STATUS"
  CURSED_MAT_STATUS MAT_STATUS = "CURSED_MAT_STATUS"
  BLESSED_MAT_STATUS MAT_STATUS = "BLESSED_MAT_STATUS"
)

type MAT_SHAPE string 
const (
  LAYER_SHAPE MAT_SHAPE = "LAYER" // ex: skin, fur?
  FEATHERS_SHAPE MAT_SHAPE = "FEATHERS" // ex: like birds
  SCALES_SHAPE MAT_SHAPE = "SCALES" // ex: like fish
  STRANDS_SHAPE MAT_SHAPE = "STRANDS" // ex: like hair
  GEOMETRIC_SHAPE MAT_SHAPE = "GEOMETRIC" // ex: celestial being
  AMORPHOUS_SHAPE MAT_SHAPE = "AMORPHOUS" // ex: blob monster, ghost
)

type MAT_STATE string
const (
  SOLID_MAT_STATE MAT_STATE = "SOLID"
  GAS_MAT_STATE  MAT_STATE = "GAS"
  LIQUID_MAT_STATE  MAT_STATE = "LIQUID"
  PLASMA_MAT_STATE  MAT_STATE = "PLASMA"
)

type Material struct {
  id MAT_NAME
  name MAT_NAME
  // TODO: think of this later
  // density int       // from 0 to 100, 0 ultra soft 100 ultra dense!
  // heat_transfer int // from 0 to 100, 0 no heat transfer, 100 a lot of heat transfer
  // conductivity int  // from 0 to 100, 0 = insulant, 100 very conductive
  // states []MAT_STATE
  // shape MAT_SHAPE 
  //statuses []MAT_STATUS
  //is_eletric_insulant bool
  //is_water_proof bool
  //is_flammable bool
  //is_explosive bool
  //reactsWith map[MAT_NAME][]ActionResultABC // maybe use another table/map for it

  // impact resistance, fractire, yield, absorption
  // compression, tension, torsion, bending, edges, texture, boil, ignite, melt point
  // densities (liquid, solid), molar_mass

  // for every organ, limb, hairs (moustache)
}

func NewMaterial() Material{
  return Material{
    id: WOOD,
    name: "regular piece of wood",
  }
}

type MATERIALS_TABLE map[MAT_NAME]Material

var MAT_TYPE_MAP MATERIALS_TABLE = map[MAT_NAME]Material{
  WOOD: Material{
    id: WOOD,
    name: WOOD,
    //density: 30, // from 0 to 100, 0 ultra soft 100 ultra dense!
    //heat_transfer: 30, // from 0 to 100, 0 no heat transfer, 100 a lot of heat transfer
    //conductivity: 30,// from 0 to 100, 0 = insulant, 100 very conductive
    //states: []MAT_STATE{SOLID_MAT_STATE},
    //shape: LAYER_SHAPE,
    //statuses: []MAT_STATUS{DRY_MAT_STATUS},
    //is_eletric_insulant: false,
    //is_water_proof: false,
    //is_flammable: true,
    //is_explosive: false,
    //reactsWith: nil, // maybe use another table/map for it
},
}

func reaction_effectiveness(s Skill, target Entity) int {
  return 10
}

func reaction(interactor, interacted Material, interaction_type InteractionType) float64 {
  efficiency, ok := reactions_map[interactor.id][interaction_type][interacted.id]
  if !ok {
    return 1
  }
  return efficiency

}
