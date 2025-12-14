package parsers

import (
	"strconv"
	"strings"

	"github.com/alexmerren/speedruncom-scraper-benchmark/internal/data"
	"github.com/bytedance/sonic"
)

type JsonIteratorParser struct{}

func (p *JsonIteratorParser) ParseBulkGame(input []byte) ([]string, error) {
	results := make([]string, 0)

	parsedInput := &data.BulkGameResponse{}
	err := sonic.Unmarshal(input, parsedInput)
	if err != nil {
		return nil, err
	}

	for _, game := range parsedInput.Data {
		result := []string{
			game.Id,
			game.Abbreviation,
			strconv.Quote(game.Names.International),
		}

		results = append(results, strings.Join(result, ","))
	}

	return results, nil
	return nil, nil
}

func (p *JsonIteratorParser) ParseGame(input []byte) ([]string, error) {
	return nil, nil
}

func (p *JsonIteratorParser) ParseLeaderboard(input []byte) ([]string, error) {
	return nil, nil
}

func (p *JsonIteratorParser) ParseUserRuns(input []byte) ([]string, error) {
	return nil, nil
}

func (p *JsonIteratorParser) GetName() string {
	return "json-iterator/go"
}
