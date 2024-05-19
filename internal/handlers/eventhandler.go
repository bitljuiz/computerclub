package handlers

import (
	"computerclub/internal/events"
	"computerclub/internal/state"
	"computerclub/internal/util"
	"fmt"
	"strconv"
)

// IncomeEventFormat defines the structure for formatting incoming events,
// including the expected body size and the behavior function to handle the event.
type IncomeEventFormat struct {
	// The expected size of the event body.
	BodySize int

	// The function to handle the event
	Behavior func(cs *state.ClubState, eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error)
}

const (
	FinallyLeft = 11 // EventID indicating that a client has finally left
	FinallySit  = 12 // EventID indicating that a client has finally been seated
	Error       = 13 // EventID indicating an error
)

var (
	// idToSize maps event IDs to the expected size of their body
	idToSize = map[events.ID]int{
		state.Came:      1,
		state.Sit:       2,
		state.IsWaiting: 1,
		state.Left:      1,
	}

	// idToBehavior maps event IDs to their corresponding behavior functions
	idToBehavior = map[events.ID]func(
		cs *state.ClubState,
		eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error){
		state.Came: func(cs *state.ClubState, eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error) {
			if !cs.IsOpened(eventTime) {
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   Error,
					OutcomeEventBody: "NotOpenYet",
				}, nil
			}
			if !state.NameIsValid(body[0]) {
				return events.OutcomeEvent{}, fmt.Errorf("invalid client name. "+
					"%sit can only contains combination letters a..z, "+
					"digits and symbols '-', '_'", util.LineSeparator())
			}
			_, ok := cs.ClientInDaClub(body[0])
			if ok {
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   Error,
					OutcomeEventBody: "YouShallNotPass",
				}, nil
			}

			cs.AddClient(state.NewClient(body[0], state.Came, -1, eventTime))

			return events.OutcomeEvent{}, nil
		},

		state.Sit: func(cs *state.ClubState, eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error) {
			client, ok := cs.ClientInDaClub(body[0])
			if !ok {
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   Error,
					OutcomeEventBody: "ClientUnknown",
				}, nil
			}

			tableIndex, _ := strconv.Atoi(body[1])

			if tableCnt := len(*cs.GetTable()); tableIndex > tableCnt {
				return events.OutcomeEvent{},
					fmt.Errorf("cannot seat on the table %d, because it is only %d in the club",
						tableIndex, tableCnt)
			}

			if !cs.JoinTheTable(client, tableIndex-1, eventTime) {
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   Error,
					OutcomeEventBody: "PlaceIsBusy",
				}, nil
			}
			return events.OutcomeEvent{}, nil
		},
		state.IsWaiting: func(cs *state.ClubState, eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error) {
			client, ok := cs.ClientInDaClub(body[0])
			if !ok {
				return events.OutcomeEvent{}, fmt.Errorf("client %s is not in da club", body[0])
			}
			queue := cs.GetWaitingQueue()
			tables := cs.GetTable()
			if len(*queue) > len(*tables) {
				_ = cs.DeleteClient(client.Name())
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   FinallyLeft,
					OutcomeEventBody: client.Name(),
				}, nil
			}
			for _, table := range *tables {
				if !table.IsUsed() {
					return events.OutcomeEvent{
						OutcomeEventTime: eventTime,
						OutcomeEventID:   Error,
						OutcomeEventBody: "ICanWaitNoLonger!",
					}, nil
				}
			}
			client.SetStatus(state.IsWaiting)
			*queue = append(*queue, client)

			return events.OutcomeEvent{}, nil
		},
		state.Left: func(cs *state.ClubState, eventTime events.TimeFormat, body []string) (events.OutcomeEvent, error) {
			leavingClient, ok := cs.ClientInDaClub(body[0])
			if !ok {
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   Error,
					OutcomeEventBody: "ClientUnknown",
				}, nil
			}
			queue := cs.GetWaitingQueue()
			switch leavingClient.Status() {
			case state.IsWaiting:
				_ = cs.DeleteClient(leavingClient.Name())
				for i, client := range *queue {
					if client.Name() == leavingClient.Name() {
						copy((*queue)[i:], (*queue)[i+1:])
						*queue = (*queue)[:len(*queue)-1]
					}
				}
				return events.OutcomeEvent{}, nil
			case state.Came:
				_ = cs.DeleteClient(leavingClient.Name())
				return events.OutcomeEvent{}, nil
			}

			cs.CalculateClient(body[0], eventTime)
			if len(*queue) > 0 {
				client := (*queue)[0]

				client.SetTable(leavingClient.TableIndex())
				client.SetTime(eventTime)
				client.SetStatus(state.Sit)
				tables := cs.GetTable()
				(*tables)[client.TableIndex()].ChangeStatus(true)
				*queue = (*queue)[1:]
				return events.OutcomeEvent{
					OutcomeEventTime: eventTime,
					OutcomeEventID:   FinallySit,
					OutcomeEventBody: client.Name() + " " + strconv.Itoa(client.TableIndex()+1),
				}, nil
			}
			return events.OutcomeEvent{}, nil
		},
	}
)

// GenerateResponseFormats generates a map of event formats for incoming events,
// mapping event IDs to their corresponding IncomeEventFormat
func GenerateResponseFormats() map[events.ID]IncomeEventFormat {
	responseFormats := make(map[events.ID]IncomeEventFormat, len(idToSize))
	for key := range idToSize {
		responseFormats[key] = IncomeEventFormat{BodySize: idToSize[key], Behavior: idToBehavior[key]}
	}
	return responseFormats
}
