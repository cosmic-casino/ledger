package types

// DONTCOVER

import "time"

const (
	// ModuleName is the name of the wta module
	ModuleName = "wta"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the wta module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the wta module
	RouterKey = ModuleName

	PrizeCollectorName = "wta_prize_collector"
	PrizeBurnerName    = "wta_prize_burner"
)

var (
	CurrentDrawEndTimeStoreKey = []byte{0x1}
	HistoricalDrawStorePrefix  = []byte("historical_draw")
	TicketsStorePrefix         = []byte("ticket")
)

// TicketsStoreKey returns the store key used to save the ticket with the given id
func TicketsStoreKey(id string) []byte {
	return append(TicketsStorePrefix, []byte(id)...)
}

// HistoricalDataStoreKey returns the store key used to save a historical data entry with the given timestamp
func HistoricalDataStoreKey(timestamp time.Time) []byte {
	return append(HistoricalDrawStorePrefix, []byte(timestamp.Format(time.RFC3339))...)
}
