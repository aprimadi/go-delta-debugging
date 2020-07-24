go-delta-debugging
==================

This library provides an implementation of the delta debugging algorithm described here: http://web2.cs.columbia.edu/~junfeng/09fa-e6998/papers/delta-debug.pdf. Code that uses this library should implement the FSM (Finite State Machine) defined in `fsm.go`.

Usage
-----

```
import (
  dd "github.com/aprimadi/go-delta-debugging"
)

// A simple FSM that becomes faulty when it contains an event "3"
type SimpleFSM struct {
  events []Event
}

// Reset the FSM state
func (f *SimpleFSM) Reset() {}

// Apply events to the FSM
func (f *SimpleFSM) Apply(events []Event) {
  f.events = events
}

// Is the FSM in valid state?
func (f *SimpleFSM) Valid() bool {
  for _, n := range f.events {
    if n == 3 {
      return false
    }
  }
  return true
}

func main() {
  fsm := SimpleFSM{}
  result := dd.DeltaDebug(fsm, []Event{1, 2, 3, 4, 5, 6, 7, 8})
  fmt.Println(result) // Print: []Event{3}
}
```
