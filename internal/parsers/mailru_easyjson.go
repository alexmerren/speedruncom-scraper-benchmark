package parsers

import (
	"strconv"
	"strings"

	"github.com/alexmerren/speedruncom-scraper-benchmark/internal/data"
	"github.com/mailru/easyjson"
)

type MailruEasyjsonParser struct{}

func (p *MailruEasyjsonParser) ParseBulkGame(input []byte) ([]string, error) {
	results := make([]string, 0)

	parsedInput := &data.BulkGameResponse{}
	err := easyjson.Unmarshal(input, parsedInput)
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

func (p *MailruEasyjsonParser) ParseGame(input []byte) ([]string, error) {
	return nil, nil
}

func (p *MailruEasyjsonParser) ParseLeaderboard(input []byte) ([]string, error) {
	return nil, nil
}

func (p *MailruEasyjsonParser) ParseUserRuns(input []byte) ([]string, error) {
	return nil, nil
}

func (p *MailruEasyjsonParser) GetName() string {
	return "mailru/easyjson"
}
