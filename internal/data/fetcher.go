package data

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func GetBulkGame() ([]byte, error) {
	url := "https://www.speedrun.com/api/v1/games?_bulk=yes&max=1000&orderby=released&direction=asc"

	response, err := getAndRetry(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func GetGame(gameId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.speedrun.com/api/v1/games/%s?embed=levels,categories,variables", gameId)

	response, err := getAndRetry(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func GetLeaderboard(gameId, categoryId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.speedrun.com/api/v1/leaderboards/%s/category/%s", gameId, categoryId)

	response, err := getAndRetry(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func GetUserRuns(userId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.speedrun.com/api/v1/runs?user=%s&max=200", userId)

	response, err := getAndRetry(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func getAndRetry(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 429 {
		time.Sleep(1_000)
		return getAndRetry(url)
	}

	slog.Info("successfully fetched data", "url", url, "status", response.StatusCode)

	return response, nil
}

func WriteToFile(data []byte, filename string) error {
	return os.WriteFile(filename, data, os.ModePerm)
}
