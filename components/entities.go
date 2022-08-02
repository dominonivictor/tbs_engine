package components

type Entity struct {
  name string
  sprite string
  combat Combat
}

func (e Entity) is_active() bool {
  return e.combat.health.is_active()
}

func NewEntity() Entity {
  combat := NewCombat()
  return Entity{
    name: "entity name",
    sprite: "E",
    combat: combat,
  }
}
