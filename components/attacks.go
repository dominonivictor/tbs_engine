package components

type Attack interface {
  attackEntity(Entity) ActionFunction
}
