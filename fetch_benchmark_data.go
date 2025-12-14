package main

import (
	"fmt"
	"os"

	"github.com/alexmerren/speedruncom-scraper-benchmark/internal/data"
)

const (
	RESOURCE_DIRECTORY       = "resources"
	SM64_GAME_ID             = "o1y9wo6q"
	SM64_16_STAR_CATEGORY_ID = "n2y55mko"
	SM64_0_STAR_CATEGORY_ID  = "xk9gg6d0"
	LARGE_USER_ID            = "qjn1wzw8"
	SMALL_USER_ID            = "8l41en2j"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%v+", err)
		os.Exit(1)
	}
}

type usecase struct {
	filename      string
	fetchFunction func() ([]byte, error)
}

func newUsecase(filename string, fetchFunction func() ([]byte, error)) usecase {
	return usecase{
		filename:      filename,
		fetchFunction: fetchFunction,
	}
}

func run() error {

	usecases := []usecase{
		newUsecase("bulk-game-data.json", func() ([]byte, error) { return data.GetBulkGame() }),
		newUsecase("large-leaderboard-data.json", func() ([]byte, error) { return data.GetLeaderboard(SM64_GAME_ID, SM64_16_STAR_CATEGORY_ID) }),
		newUsecase("small-leaderboard-data.json", func() ([]byte, error) { return data.GetLeaderboard(SM64_GAME_ID, SM64_0_STAR_CATEGORY_ID) }),
		newUsecase("large-user-data.json", func() ([]byte, error) { return data.GetUserRuns(LARGE_USER_ID) }),
		newUsecase("small-user-data.json", func() ([]byte, error) { return data.GetUserRuns(SMALL_USER_ID) }),
		newUsecase("game-data.json", func() ([]byte, error) { return data.GetGame(SM64_GAME_ID) }),
	}

	for _, usecase := range usecases {
		usecaseData, err := usecase.fetchFunction()
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("%s/%s", RESOURCE_DIRECTORY, usecase.filename)

		err = data.WriteToFile(usecaseData, filename)
		if err != nil {
			return err
		}
	}

	return nil
}
