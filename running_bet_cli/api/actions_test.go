package api_test

import (
	"bytes"
	"testing"

	"github.com/bedminer1/running_bet/api"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		testName string
		myScore  float32
		herScore float32
		scheme   [][]float32
		expRes   api.Record
	}{
		{
			testName: "tieWinNoScheme",
			myScore: float32(10),
			herScore: float32(10),
			scheme: [][]float32{},
			expRes: api.Record{
				WeekID: 1,
				MyScore: float32(10),
				HerScore: float32(10),
				NeededScore: float32(10),
				WinForMe: false,
				MyPoints: 0,
				HerPoints: 1,
				Scheme: [][]float32{},
			},
		},
		{
			testName: "sheWinsAgain",
			myScore:  30,
			herScore: 20,
			scheme:   [][]float32{{10, 2}, {20, 1.5}, {30, 1}},
			expRes: api.Record{
				WeekID: 2,
				MyScore:     30,
				HerScore:    20,
				NeededScore: 35.0, // (10*2 + 10*1.5 + 0*1)
				WinForMe:    false,
				MyPoints:    0,
				HerPoints:   2,
				Scheme:      [][]float32{{10, 2}, {20, 1.5}, {30, 1}},
			},
		},
		{
			testName: "iWin",
			myScore:  15,
			herScore: 7,
			scheme:   [][]float32{{7, 2}, {12, 1}},
			expRes: api.Record{
				WeekID: 3,
				MyScore:     15,
				HerScore:    7,
				NeededScore: 14.0, // (7*2 + 0*1)
				WinForMe:    true,
				MyPoints:    1,
				HerPoints:   2,
				Scheme:      [][]float32{{7, 2}, {12, 1}},
			},
		},
	}

	db := &api.Records{}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := db.Add(tc.myScore, tc.herScore, tc.scheme)
			if err != nil {
				t.Errorf("Unexpected error with db.Add: %q", err)
			}

			// Get the last entry in db (the one just added)
			res := (*db)[len(*db)-1]

			if res.WeekID != tc.expRes.WeekID {
				t.Errorf("WeekID mismatch: expected %d, got %d", tc.expRes.WeekID, res.WeekID)
			}
			if res.NeededScore != tc.expRes.NeededScore {
				t.Errorf("NeededScore mismatch: expected %.2f, got %.2f", tc.expRes.NeededScore, res.NeededScore)
			}
			if res.WinForMe != tc.expRes.WinForMe {
				t.Errorf("WinForMe mismatch: expected %t, got %t", tc.expRes.WinForMe, res.WinForMe)
			}
			if res.MyPoints != tc.expRes.MyPoints {
				t.Errorf("MyPoints mismatch: expected %d, got %d", tc.expRes.MyPoints, res.MyPoints)
			}
			if res.HerPoints != tc.expRes.HerPoints {
				t.Errorf("HerPoints mismatch: expected %d, got %d", tc.expRes.HerPoints, res.HerPoints)
			}
		})
	}
}

func TestGetTxt(t *testing.T) {
	expData := api.Record{
		MyScore:  float32(36.85),
		WinForMe: true,
	}

	db := &api.Records{}
	fileDirectory := "../local_storage/"
	fileName := "record"
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

func TestGetJSON(t *testing.T) {
	expData := api.Record{
		MyScore:  float32(36.85),
		WinForMe: true,
	}

	db := &api.Records{}
	fileDirectory := "../local_storage/"
	fileName := "record"
	fileType := "json"

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
	fileType := "json"
	if err := db.Get(fileDirectory, fileName, fileType); err != nil {
		t.Error("Unexpected error getting data")
	}

	if err := db.Save(fileDirectory, fileName); err != nil {
		t.Errorf("Unexpected error saving file: %q", err)
	}
}

func TestList(t *testing.T) {
	db := &api.Records{}

	expOut := `WeekID  |MyScore  |HerScore  |NeededScore  |WinForMe   |MyPoints   |HerPoints  |Scheme
------  |-------  |--------  |-----------  |---------  |---------  |---------  |------
1       |36.85    |9.07      |18.14        |true       |1          |0          |[[1000 2]]
`

	// fetch test data
	fileDirectory := "../local_storage/"
	fileName := "record_test"
	fileType := "json"
	if err := db.Get(fileDirectory, fileName, fileType); err != nil {
		t.Error("Unexpected error getting data")
	}

	var out bytes.Buffer
	db.List(&out)

	if expOut != out.String() {
		t.Errorf("Print output mismatch: expected %s, got %s", expOut, out.String())
	}

}