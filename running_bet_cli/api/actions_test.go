package api_test

import (
	"fmt"
	"testing"

	"github.com/bedminer1/running_bet/api"
)

func TestGetTxt(t *testing.T) {
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

func TestSave(t *testing.T) {
	db := &api.Records{}
	fileDirectory := "../local_storage/"
	fileName := "record"
	fileType := "txt"
	if err := db.Get(fileDirectory, fileName, fileType); err != nil {
		t.Error("Unexpected error getting data")
	}
	for i := range *db {
		r := &(*db)[i]
		if r.WeekID < 11 {
			r.Scheme = [][]float32{{1000, 2}}
		} else if r.WeekID < 22 {
			r.Scheme = [][]float32{{1000, 1.5}}
		} else if r.WeekID < 27 {
			r.Scheme = [][]float32{}
		} else {
			r.Scheme = [][]float32{{5, 2}, {10, 1.5}}
		}
	}

	fmt.Println(*db)
	if err := db.Save(fileDirectory, fileName); err != nil {
		t.Errorf("Unexpected error saving file: %q", err)
	}
}