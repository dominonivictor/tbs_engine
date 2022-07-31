package components

type MAT_SHAPE string 
const (
  LAYER MAT_SHAPE = "LAYER" // ex: skin, fur?
  FEATHERS MAT_SHAPE = "FEATHERS" // ex: like birds
  SCALES MAT_SHAPE = "SCALES" // ex: like fish
  STRANDS MAT_SHAPE = "STRANDS" // ex: like hair
)
type MAT_STATE string
const (
  SOLID MAT_STATE = "SOLID"
  GAS MAT_STATE = "GAS"
  LIQUID MAT_STATE = "LIQUID"
  PLASMA MAT_STATE = "PLASMA"
)

type MaterialInteractable interface {
  interactWithMaterial(m MaterialTemplate)
}
type MaterialTemplate struct {
  name MAT_NAME
  density int // from 0 to 100, 0 not soft, 100 ultra hard
  heat_transfer int // from 0 to 100, 0 no heat transfer, 100 a lot of heat transfer
  conductivity int // from 0 to 100, 0 = insulant, 100 very conductive
  state MAT_STATE
  shape MAT_SHAPE
  is_eletric_insulant bool
  is_water_proof bool
  is_flammable bool
  is_explosive bool
  reactsWith MaterialInteractable // maybe use another table/map for it

  // impact resistance, fractire, yield, absorption
  // compression, tension, torsion, bending, edges, texture, boil, ignite, melt point
  // densities (liquid, solid), molar_mass

  //nail Nail
  //claw Claw
  //scale Scale
  //...
  // for every organ, limb, hairs (moustache)
}

type MAT_NAME string
const (
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
type MAT_TYPE map[MAT_NAME]MaterialTemplate

var MAT_TYPE_MAP MAT_TYPE = map[MAT_NAME]MaterialTemplate{
  WOOD: MaterialTemplate{name: WOOD, density: 1, },
}

type Material struct {
  type_id MAT_NAME

}

func (m Material) interactWithAction(a Action) {

}

func (m Material) interactWithAttack(a Attack) {

}

func (m1 Material) interactWithMaterial(m2 Material) {

}
