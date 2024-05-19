## handlers

The `handlers` package contains an implementation of the arrival handler
events, which defines behavior depending on `IncomeEventID`
events accordingly.

The file `eventhandler.go` provides the factory implementation
handlers, each of which has its own `IncomeEventID`.
When calling `GenerateResponseFormats` all possible
handlers for conditions-known `ID`

- `IncomeEventFormat` represents the handler storage format for
  specific event depending on its `IncomeEventID`