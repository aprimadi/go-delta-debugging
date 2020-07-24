package deltadebug

import (
  "reflect"
  "testing"
)

type TestFSM struct {
  numbers map[int]struct{}
  faulty map[int]struct{}
}

func NewTestFSM(faulty []int) *TestFSM {
  fsm := new(TestFSM)
  fsm.numbers = make(map[int]struct{})
  fsm.faulty = make(map[int]struct{})
  for _, n := range faulty {
    fsm.faulty[n] = struct{}{}
  }
  return fsm
}

func (f *TestFSM) Reset() {
  f.numbers = make(map[int]struct{})
}

func (f *TestFSM) Apply(events []Event) {
  for _, op := range events {
    n := op.(int)
    f.numbers[n] = struct{}{}
  }
}

func (f *TestFSM) Valid() bool {
  for n, _ := range f.faulty {
    if _, ok := f.numbers[n]; !ok {
      return true
    }
  }
  return false
}

func TestSimple(t *testing.T) {
  fsm := NewTestFSM([]int{3, 6})
  valid := [][]Event{
    []Event{1, 2, 3, 4},
    []Event{5, 6, 7, 8},
  }
  invalid := [][]Event{
    []Event{1, 2, 3, 4, 5, 6, 7, 8},
    []Event{3, 6},
    []Event{3, 5, 6, 7, 8},
    []Event{3, 4, 5, 6, 7, 8},
  }
  for _, e := range valid {
    fsm.Reset()
    fsm.Apply(e)
    if !fsm.Valid() {
      t.Errorf("FSM should be valid: %v", e)
    }
  }
  for _, e := range invalid {
    fsm.Reset()
    fsm.Apply(e)
    if fsm.Valid() {
      t.Errorf("FSM should be invalid: %v", e)
    }
  }

  result := DeltaDebug(fsm, []Event{1, 2, 3, 4, 5, 6, 7, 8})
  if !reflect.DeepEqual(result, []Event{3, 6}) {
    t.Errorf("Want: %v, got: %v", []Event{3, 6}, result)
  }
}
