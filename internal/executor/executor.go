package executor

import (
	"computerclub/internal/events"
	"computerclub/internal/handlers"
	"computerclub/internal/reader"
	"computerclub/internal/state"
	"computerclub/internal/util"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// Execute processes the events of the computer club and if there is some problem
// it would return line, where it was detected and error message
func Execute(filepath string) (string, error) {
	r, err := reader.NewFileEventReader(filepath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	var tableCount int
	if firstLine, err := r.ReadLine(); err == nil {
		tableCount, err = strconv.Atoi(firstLine)
		if err != nil {
			return firstLine, err
		}
		if tableCount < 1 {
			return firstLine, fmt.Errorf("table count must be positive integer value")
		}
	} else {
		return firstLine, err
	}

	var openTime, closeTime events.TimeFormat
	if secondLine, err := r.ReadLine(); err == nil {
		times := strings.Split(secondLine, " ")
		if len(times) != 2 {
			return secondLine, fmt.Errorf("second line must contains only two tokens.%s"+
				"Time when club opens in format XX:XX and time when club closes in format XX:XX%s."+
				"Tokens must be separated by single space", util.LineSeparator(), util.LineSeparator())
		}
		if openTime, err = events.ParseTime(times[0]); err != nil {
			return secondLine, fmt.Errorf("invalid format of opening time. Valid input looks like XX:XX, for example: 08:48.%s", util.LineSeparator())
		}
		if closeTime, err = events.ParseTime(times[1]); err != nil {
			return secondLine, fmt.Errorf("invalid format of closing time. Valid input looks like XX:XX, for example: 08:48.%s", util.LineSeparator())
		}
		if !openTime.EarlierThan(closeTime) {
			return secondLine, fmt.Errorf("opening time must be earlier than closing time")
		}
	} else {
		return secondLine, err
	}

	var costPerHour int
	if thirdLine, err := r.ReadLine(); err == nil {
		costPerHour, err = strconv.Atoi(thirdLine)
		if err != nil {
			return thirdLine, err
		}
		if costPerHour < 1 {
			return thirdLine, fmt.Errorf("cost must be integer positive number")
		}
	} else {
		return thirdLine, err
	}

	handler := handlers.GenerateResponseFormats()
	club := state.NewClubState(tableCount, openTime, closeTime, uint(costPerHour))
	processedEvents := make([]events.Event, 0)
	var lastEventTime events.TimeFormat

	for {
		event, err := r.ReadEvent()
		if err != nil {
			if err == io.EOF {
				break
			}
			return event.String(), err
		}
		if event.IncomeEventTime.EarlierThan(lastEventTime) {
			return event.String(), fmt.Errorf("each event should be not earlier than previous%sPrevious: %v",
				util.LineSeparator(), lastEventTime)
		}
		lastEventTime = event.IncomeEventTime
		processedEvents = append(processedEvents, event)

		if executor, ok := handler[event.IncomeEventID]; ok {
			body := strings.SplitN(event.IncomeEventBody, " ", executor.BodySize)
			if len(body) != executor.BodySize {
				return event.String(), fmt.Errorf("event %v must contains %d tokens, but found %d",
					event, executor.BodySize, len(body))
			}
			outcomeEvent, err := executor.Behavior(club, event.IncomeEventTime, body)
			if err != nil {
				return event.String(), err
			}
			if outcomeEvent.OutcomeEventBody != "" {
				processedEvents = append(processedEvents, outcomeEvent)
			}
		} else {
			return event.String(), fmt.Errorf("cannot find event with id %d", event.IncomeEventID)
		}
	}

	leftClient := club.Clients()
	sort.Strings(leftClient)
	for _, clientName := range leftClient {
		if client, _ := club.ClientInDaClub(clientName); client.Status() != state.Sit {
			_ = club.DeleteClient(clientName)
		} else {
			club.CalculateClient(clientName, closeTime)
		}
		processedEvents = append(processedEvents,
			events.OutcomeEvent{
				OutcomeEventTime: closeTime,
				OutcomeEventID:   handlers.FinallyLeft,
				OutcomeEventBody: clientName,
			},
		)
	}
	fmt.Println(openTime)
	for _, event := range processedEvents {
		fmt.Println(event)
	}
	fmt.Println(closeTime)
	for tableIndex, table := range *club.GetTable() {
		fmt.Printf("%d %d %v%s", tableIndex+1, table.Earnings(), table.UsageTime(), util.LineSeparator())
	}
	return "", nil
}
