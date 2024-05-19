## events

The `events` package contains the formation of events happen in the computer club

The file `event.go` implements the general event interface `Event`
and two of its implementations
- `IncomeEvent` - format of incoming events
- `OutcomeEvent` - formatting of outgoing events

The `factory.go` file implements the assembly concept of
`IncomeEvent` from strings. It is assumed that if the input data format
changes, and campaigns collect `IncomeEvent` in some other way, then
the implementation will be supplemented in this file

- `EventFactory` is a large purpose structure.
  which is the creation of IncomeEvent

The file `timeformat.go` contains the time form that is used.
throughout the project

- `TimeFormat` structure, which stores time in XX:XX format.