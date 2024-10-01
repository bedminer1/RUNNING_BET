package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	WeekID      int         `json:"weekID"`
	MyScore     float32     `json:"my_score"`
	HerScore    float32     `json:"her_score"`
	NeededScore float32     `json:"needed_score"`
	WinForMe    bool        `json:"win_for_me"`
	MyPoints    int         `json:"my_points"`
	HerPoints   int         `json:"her_points"`
	Scheme      [][]float32 `json:"scheme"`
}

type Records []Record

func (db *Records) Add(myScore, herScore float32, scheme [][]float32) error {
	neededScore := calculateScore(herScore, scheme)
	var myPoints, herPoints int
	if len(*db) > 0 {
		myPoints = (*db)[len(*db)-1].MyPoints
		herPoints = (*db)[len(*db)-1].HerPoints
	}

	if myScore > neededScore {
		myPoints++
	} else {
		herPoints++
	}

	r := Record{
		WeekID:      len(*db)+1,
		MyScore:     myScore,
		HerScore:    herScore,
		NeededScore: neededScore,
		WinForMe:    myScore > neededScore,
		MyPoints:    myPoints,
		HerPoints:   herPoints,
		Scheme:      scheme,
	}
	*db = append(*db, r)
	return nil
}

func (db *Records) Get(fileDirectory, fileName, fileType string) error {
	filePath := fmt.Sprint(fileDirectory, fileName, ".", fileType)
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	switch fileType {
	case "txt":
		*db, err = parseTxt(f)
	case "json":
		*db, err = parseJSON(f)
	default:
		return fmt.Errorf("fileType not supported")
	}
	if err != nil {
		return err
	}
	return nil
}

func (db *Records) Save(fileDirectory, fileName string) error {
	filePath := fmt.Sprintf("%s%s.json",fileDirectory, fileName)
	js, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, js, 0644)
}
