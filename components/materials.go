package components

import (
  "fmt"
)

var reactions_map map[MAT_ID]map[InteractionType]map[MAT_ID]float64 = map[MAT_ID]map[InteractionType]map[MAT_ID]float64{
  WOOD: map[InteractionType]map[MAT_ID]float64{
    ATK_INTERACTION: map[MAT_ID]float64{
      WOOD: 1,
    },
  },
  FIRE: map[InteractionType]map[MAT_ID]float64{
    ATK_INTERACTION: map[MAT_ID]float64{
      WOOD: 2,
    },
  },
}

type MAT_ID string
const (
  FIRE MAT_ID = "FIRE"
  WATER MAT_ID = "WATER"
  EARTH MAT_ID = "EARTH"
  AIR MAT_ID = "AIR"
  LIGHTNING MAT_ID = "LIGHTNING"
  WOOD MAT_ID = "WOOD"
  CLOTH MAT_ID = "CLOTH"
  BONE MAT_ID = "BONE"
  FAT MAT_ID = "FAT"
  FEATHER MAT_ID = "FEATHER"
  STONE MAT_ID = "STONE"
  MUSCLE MAT_ID = "MUSCLE"
  VAPOR MAT_ID = "VAPOR"
  RUBBER MAT_ID = "RUBBER"
  HAIR MAT_ID = "HAIR"
  METAL MAT_ID = "METAL"
  SCALE MAT_ID = "SCALE"
  LEATHER MAT_ID = "LEATHER"
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
  id MAT_ID
  name string
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
  //reactsWith map[MAT_ID][]ActionResultABC // maybe use another table/map for it

  // impact resistance, fractire, yield, absorption
  // compression, tension, torsion, bending, edges, texture, boil, ignite, melt point
  // densities (liquid, solid), molar_mass

  // for every organ, limb, hairs (moustache)
}

func NewMaterialFromCSV(args ...string) Material {
  return Material{
    id: MAT_ID(args[0]),
    name: args[1],

  }

}

func NewMaterial(name MAT_ID) Material{
  var m Material
  switch name {
    case WOOD:
      m = Material{
            id: WOOD,
            name: "regular piece of wood",
          }
    default:
      m = Material{
        id: FIRE,
        name: "the fiery fire",
      } 
 
  }
  return m
}

type MATERIALS_TABLE map[MAT_ID]Material

var MAT_TYPE_MAP MATERIALS_TABLE = map[MAT_ID]Material{
  WOOD: Material{
    id: WOOD,
    name: "woody",
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
  fmt.Printf("reactions map: %+v \n", reactions_map)
  fmt.Printf("map[%s][%s][%s]", interactor.id, interaction_type, interacted.id)
  efficiency, ok := reactions_map[interactor.id][interaction_type][interacted.id]
  if !ok {
    return 1
  }
  return efficiency

}
