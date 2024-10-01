package api

import (
	"fmt"
	"os"
)

type Record struct {
	WeekID      int
	MyScore     float32
	HerScore    float32
	NeededScore float32
	WinForMe    bool
	MyPoints    int
	HerPoints   int
	Scheme      [][]float32
}

type Records []Record

func (db *Records) Add(myScore, herScore float32, scheme [][]float32) error {
	neededScore := calculateScore(herScore, scheme)
	myPoints := (*db)[len(*db)-1].MyPoints + 1
	herPoints := (*db)[len(*db)-1].HerPoints + 1


	if myScore > neededScore {
		myPoints++
	} else {
		herPoints++
	}
	
	r := Record{
		WeekID: len(*db),
		MyScore: myScore,
		HerScore: herScore,
		NeededScore: neededScore,
		WinForMe: myScore > neededScore,
		MyPoints: myPoints,
		HerPoints: herPoints,
		Scheme: scheme,
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
	default:
		return fmt.Errorf("fileType not supported")
	}
	if err != nil {
		return err
	}
	return nil
}

func (db *Records) Save(fileName string) error {
	return nil
}