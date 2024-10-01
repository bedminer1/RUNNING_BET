package api_test

import (
	"testing"

	"github.com/bedminer1/running_bet/api"
)

func TestGet(t *testing.T) {
	db := &api.Records{}
	fileName := "../local_storage/record.txt"
	fileType := "txt"
	if err := db.Get(fileName, fileType); err != nil {
		t.Errorf("Unexpected error: %q", err)
	}
}