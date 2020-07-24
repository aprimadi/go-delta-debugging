package deltadebug

type Event interface{}

type FSM interface {
  // Reset the FSM state
  Reset()

  // Apply a set of operations to the FSM
  Apply(events []Event)

  // Test whether the FSM is in a valid state
  Valid() bool
}
