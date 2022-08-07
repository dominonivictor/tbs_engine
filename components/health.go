package components

import (
  "fmt"
)

type HealthMng struct {
  // for now it will be very simple, hp, max hp, thats it
  core Health
}

func NewHealthMng(_ *LoadedData, _ Breed) HealthMng {
  return HealthMng{
    core: NewHealth(),
  } 
}

func (h HealthMng) is_active() bool {
  fmt.Printf("checking if its active h.core.hp=%d > 0 ?=%t\n", h.core.hp, h.core.hp > 0)
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

func (h HealthMng) get_bonus() int {
  return 0
}

