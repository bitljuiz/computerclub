## events

Пакет `events` содержит в себе реализацию формата событий,
которые происходят в компьютерном клубе

В файле `event.go` реализован общий интерфейс события `Event` 
и две его реализации
- `IncomeEvent` - формат входящих событий
- `OutcomeEvent` - формат исходящих событий

В файле `factory.go` реализована концепция сборки
`IncomeEvent` из строки. Предполагается, что если формат входных данных 
поменяется и придется собирать `IncomeEvent` как-либо иначе, то 
реализация будет дополнена именно в этом файле

- `EventFactory` представляет собой структуру, единственное предназначение
которой является создание `IncomeEvent`

В файле `timeformat.go` находится реализация формата времени, которое используется 
во всем проекте

- `TimeFormat` структура, в котором хранится время, в формате XX:XX.