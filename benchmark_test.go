package main_test

import (
	"os"
	"testing"

	"github.com/alexmerren/speedruncom-scraper-benchmark/internal/parsers"
)

var subjects = []parsers.Parser{
	&parsers.BugerJsonparserParser{},
	&parsers.BytedanceSonicParser{},
	&parsers.JsonIteratorParser{},
	&parsers.MailruEasyjsonParser{},
}

func Benchmark_BulkGameData(b *testing.B) {
	// Given
	data, err := loadBulkGameData()
	check(err)

	for _, subject := range subjects {
		b.Run(subject.GetName(), func(b *testing.B) {
			for b.Loop() {
				subject.ParseBulkGame(data)
			}
		})
	}
}

func loadBulkGameData() ([]byte, error) {
	return os.ReadFile("resources/bulk-game-data.json")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
