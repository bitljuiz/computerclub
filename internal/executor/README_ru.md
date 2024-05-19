## executor

В пакете `executor` хранится только функция
`Execute`, которая считывает из файла все данные о 
компьютерном клубе и различные события, обрабатывает их
и если никаких ошибок формата в процессе не было получено, то
выводит в консоль:
- обработанные входящие события
- исходящие события
- время открытия и закрытия клуба
- прибыль и суммарное время работы каждого стола

Если в процессе работы были выявлены ошибки формата, то
`Execute` вернет первую строчку, где была получена ошибка, соответственно
описание ошибки и прекратит исполнение