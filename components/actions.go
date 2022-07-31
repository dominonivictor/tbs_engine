package components

type Action interface{
  useOn(Entity) ActionFunction
}

type ActionFunction interface {
  act()
}
