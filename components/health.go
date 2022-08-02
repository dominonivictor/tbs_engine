package components

type HealthMng struct {
  // for now it will be very simple, hp, max hp, thats it
  core Health
}

func NewHealthMng() HealthMng {
  return HealthMng{
    core: NewHealth(),
  } 
}

func (h HealthMng) is_active() bool {
  return h.core.hp > 0
}

type Health struct {
  hp int
  max_hp int
}

func NewHealth() Health {
  return Health{
    hp: 10,
    max_hp: 10,
  }
}

func (h HealthMng) get_bonus() float64 {
  return 30
}

