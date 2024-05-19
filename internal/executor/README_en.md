## executor

The `executor` package stores only the function
`Execute`, which reads all data about
computer club and various events, processes them
and if no format errors were received during the process, then
outputs to the console:
- processed income events
- outcome events
- opening and closing times of the club
- profit and total operating time of each table

If format errors were identified during the work, then
`Execute` will return the first line where the error was received, respectively
description of the error and stops execution