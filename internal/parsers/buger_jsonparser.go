package parsers

import (
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

type BugerJsonparserParser struct{}

func (p *BugerJsonparserParser) ParseBulkGame(input []byte) ([]string, error) {
	results := make([]string, 0)

	jsonparser.ArrayEach(input, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		gameId, _ := jsonparser.GetString(value, "id")
		gameName, _ := jsonparser.GetString(value, "names", "international")
		gameAbbreviation, _ := jsonparser.GetString(value, "abbreviation")

		result := []string{
			gameId,
			gameAbbreviation,
			strconv.Quote(gameName),
		}

		results = append(results, strings.Join(result, ","))
	}, "data")

	return results, nil
}

func (p *BugerJsonparserParser) ParseGame(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BugerJsonparserParser) ParseLeaderboard(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BugerJsonparserParser) ParseUserRuns(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BugerJsonparserParser) GetName() string {
	return "buger/jsonparser"
}
