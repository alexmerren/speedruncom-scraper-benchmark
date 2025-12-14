package parsers

// A Parser creates any number of CSV records from a given API response. This is intended to simulate the parsing behaviour of speedruncom-scraper processing.
type Parser interface {

	// Parse the response from [data.GetBulkGame] into a CSV record for each game.
	ParseBulkGame(input []byte) ([]string, error)

	// Parse the response from [data.GetGame] into a single CSV record.
	ParseGame(input []byte) ([]string, error)

	// Parse the response from [data.GetLeaderboard] into a CSV record for each leaderboard entry.
	ParseLeaderboard(input []byte) ([]string, error)

	// Parse the response from [data.GetUserRuns] into a CSV record for each run.
	ParseUserRuns(input []byte) ([]string, error)

	GetName() string
}
