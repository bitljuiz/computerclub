## reader

В пакете `reader` описан считыватель `FileEventReader`, 
реализующий интерфейс `EventReader`

В файле `reader.go` реализованы
- `EventReader` интерфейс, описывающий, как должен выглядеть считыватель для `IncomeEvent`
- `FileEventReader` конкретная реализация `EventReader`, предназначенная, для считывания из файла
