package parsers

import (
	"strconv"
	"strings"

	"github.com/alexmerren/speedruncom-scraper-benchmark/internal/data"
	"github.com/bytedance/sonic"
)

type BytedanceSonicParser struct{}

func (p *BytedanceSonicParser) ParseBulkGame(input []byte) ([]string, error) {
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
}

func (p *BytedanceSonicParser) ParseGame(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BytedanceSonicParser) ParseLeaderboard(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BytedanceSonicParser) ParseUserRuns(input []byte) ([]string, error) {
	return nil, nil
}

func (p *BytedanceSonicParser) GetName() string {
	return "bytedance/sonic"
}
