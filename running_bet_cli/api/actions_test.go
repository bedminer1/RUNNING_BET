package api_test

import (
	"testing"

	"github.com/bedminer1/running_bet/api"
)

func TestGet(t *testing.T) {
	expData := api.Record{
		MyScore:     float32(36.85),
		WinForMe:    true,
	}

	db := &api.Records{}
	fileDirectory := "../local_storage/"
	fileName := "record_test"
	fileType := "txt"
	if err := db.Get(fileDirectory, fileName, fileType); err != nil {
		t.Errorf("Unexpected error: %q", err)
	}
	res := (*db)[0]

	if expData.MyScore != res.MyScore {
		t.Errorf("Expected score: %f, Recorded Score: %f", expData.MyScore, res.MyScore)
	}

	if expData.WinForMe != res.WinForMe {
		t.Errorf("Expected WinForMe: %t, Recorded WinForMe: %t", expData.WinForMe, res.WinForMe)
	}
}
