## handlers

В пакете `handlers` находится реализация обработчика поступивших
событий, который определяет поведение в зависимости от `IncomeEventID`
события соответственно.

В файле `eventhandler.go` представлена реализация фабрики
обработчиков, каждому из которых соответствует свой `IncomeEventID`. 
При вызове `GenerateResponseFormats` будут сгенерированы все возможные 
обработчики для известных по условию `ID`

- `IncomeEventFormat` представляет собой формат хранения обработчика для 
конкретного события в зависимости от его `IncomeEventID`