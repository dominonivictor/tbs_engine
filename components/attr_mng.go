package components


type AttrMng struct {
  list []Attr

}

func NewAttrMng(_ *LoadedData, _ Breed) AttrMng {
  return AttrMng{
    list: []Attr{
      NewAttr(),
    },
  }
}

func (a AttrMng) get_bonus(s *Skill) int {
  return 0
}

type Attr struct {
  name string
}

func NewAttr() Attr {
  return Attr{
    name: "attr name",
  }
}
