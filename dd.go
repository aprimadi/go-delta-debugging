package deltadebug

// Returns the minimum set of operations that caused FSM to be in invalid state
func DeltaDebug(fsm FSM, events []Event) []Event {
  return dd(fsm, events, []Event{})
}

func dd(fsm FSM, events []Event, remaining []Event) []Event {
  n := len(events)

  // Termination condition
  if n == 0 {
    return []Event{}
  }
  if n == 1 {
    if valid(fsm, append(events, remaining...)) {
      return []Event{}
    } else {
      return events
    }
  }

  // Split the operations in two
  split := (n + 1) / 2
  evs1 := append([]Event{}, events[:split]...)
  evs2 := append([]Event{}, events[split:]...)

  if !valid(fsm, append(evs1, remaining...)) {
    return dd(fsm, evs1, remaining)
  }
  if !valid(fsm, append(evs2, remaining...)) {
    return dd(fsm, evs2, remaining)
  }

  // Do interposition
  ievs1 := dd(fsm, evs1, evs2)
  ievs2 := dd(fsm, evs2, evs1)
  return append(ievs1, ievs2...)
}

func valid(fsm FSM, events []Event) bool {
  fsm.Reset()
  fsm.Apply(events)
  return fsm.Valid()
}
