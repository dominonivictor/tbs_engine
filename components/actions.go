package components

import (
  "fmt"
)

type ActionABC interface{
  act(...Entity) ActionResultABC
}

type ActionResultABC interface {
  show_results()
  String() string
}

type ActionResult struct {
  text string
}
func (a ActionResult) show_results() {
  fmt.Println(a.text)
}
func (a ActionResult) String() string {
  return fmt.Sprint(a.text)
}
