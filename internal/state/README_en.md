## state

The `state` package implements entities that, collectively,
represent the current state of the computer club

The `client.go` file contains the presentation method
client data, as well as all changes that may occur to it

- `Client` is a structure that stores information about the client: name,
  current status (`Came`, `Sit`, `IsWaiting` and `Left`)

The file `table.go` implements the representation of table data in the computer club

- `Table` is a structure that stores data about the table: whether it is currently occupied,
  how much was earned per day, as well as the total time of use (in XX:XX format)

The file `state.go` fully describes the implementation of the computer club state. This is a fundamental part of the project,
in which everything that changes in the state of tables and clients is stored. Data about the work of the club itself is also stored:
opening and closing times, cost per hour, number of tables

- `ClubState` is an implementation of an abstract representation of the current state of the club. All incoming requests
  one way or another will receive information from `ClubState`