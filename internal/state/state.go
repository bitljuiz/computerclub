package state

import (
	"computerclub/internal/events"
	"fmt"
)

// ClubState represents the state of the computer club
type ClubState struct {
	clients      map[string]*Client // A map of clients in the club. Key is client name, value is Client
	tables       []Table            // A slice of tables in the club
	waitingQueue []*Client          // A queue of clients waiting for a table
	openingTime  events.TimeFormat  // The opening time of the club
	closingTime  events.TimeFormat  // The closing time of the club
	costPerHour  uint               // The cost per hour of using a table
}

// AddClient adds a new client to the club
func (cs *ClubState) AddClient(client Client) {
	cs.clients[client.name] = &client
}

// DeleteClient removes a client from the club by their name
// Returns an error if the client is not found
func (cs *ClubState) DeleteClient(clientName string) error {
	_, ok := cs.clients[clientName]
	if !ok {
		return fmt.Errorf("cannot found client with name %s", clientName)
	}
	delete(cs.clients, clientName)
	return nil
}

// GetWaitingQueue returns a pointer to the waiting queue of clients
func (cs *ClubState) GetWaitingQueue() *[]*Client {
	return &cs.waitingQueue
}

// IsOpened checks if the club is open at the given time
func (cs ClubState) IsOpened(time events.TimeFormat) bool {
	return cs.openingTime.EarlierThan(time) && time.EarlierThan(cs.closingTime)
}

// ClientInDaClub checks if a client is in the club by their name
// Returns the client and a boolean indicating if the client is in da club
func (cs ClubState) ClientInDaClub(clientName string) (*Client, bool) {
	client, ok := cs.clients[clientName]
	return client, ok
}

// CalculateClient calculates the total cost and usage time for a client and updates the table's usage statistics
// Returns false if the client is not found
func (cs *ClubState) CalculateClient(clientName string, endTime events.TimeFormat) bool {
	client, ok := cs.ClientInDaClub(clientName)
	if !ok {
		return false
	}

	table := &cs.tables[client.tableIndex]
	diff, _ := client.startTime.Difference(endTime)
	table.usageTime += diff
	hours := uint(diff / 60)
	if diff%60 != 0 {
		hours += 1
	}
	table.earnings += hours * cs.costPerHour
	table.isUsed = false
	delete(cs.clients, clientName)
	return true
}

// JoinTheTable assigns a client to a table if the table is not currently in use
// Returns false if the table is already in use
func (cs *ClubState) JoinTheTable(client *Client, tableIndex int, time events.TimeFormat) bool {
	table := &cs.tables[tableIndex]
	if table.isUsed {
		return false
	}
	table.isUsed = true

	if client.status == Sit {
		_ = cs.CalculateClient(client.name, time)
		cs.clients[client.name] = client
	}
	client.tableIndex = tableIndex
	client.startTime = time
	client.status = Sit
	return true
}

// GetTable returns a pointer to the slice of tables in the club
func (cs *ClubState) GetTable() *[]Table {
	return &cs.tables
}

// Clients returns a slice of client names currently in the club
func (cs *ClubState) Clients() []string {
	keys := make([]string, 0)
	for key := range cs.clients {
		keys = append(keys, key)
	}
	return keys
}

// NewClubState creates a new ClubState
func NewClubState(tablesCount int, openingTime events.TimeFormat, closingTime events.TimeFormat,
	costPerHour uint) *ClubState {
	clients := make(map[string]*Client, tablesCount)

	tables := make([]Table, tablesCount)

	for i := range tables {
		tables[i] = Table{isUsed: false, earnings: 0}
	}

	waitingQueue := make([]*Client, 0)

	return &ClubState{clients: clients, tables: tables, waitingQueue: waitingQueue, openingTime: openingTime, closingTime: closingTime,
		costPerHour: costPerHour}
}
