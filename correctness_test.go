package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseBulkGame(t *testing.T) {
	// Given
	data := []byte(`{"data":[{"id":"j1n8nj91","names":{"international":"Computer Space","japanese":null},"abbreviation":"computer_space","weblink":"http://www.speedrun.com/computer_space"}],"pagination":{"offset":0,"max":1000,"size":1,"links":[{"rel":"next","uri":"http://www.speedrun.com/api/v1/games?_bulk=yes&max=1000&orderby=released&direction=asc&offset=1000"}]}}`)
	expectedResult := `j1n8nj91,computer_space,"Computer Space"`

	for _, subject := range subjects {
		t.Run(subject.GetName(), func(t *testing.T) {
			// When
			result, err := subject.ParseBulkGame(data)

			// Then
			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, 1, len(result))
			assert.Equal(t, expectedResult, result[0])
		})
	}
}
