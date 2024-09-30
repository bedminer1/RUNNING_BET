package api

import "os"

type Record struct {
	weekID      int
	myScore     float32
	herScore    float32
	neededScore float32
	winForMe    bool
	myPoints    int
	herPoints   int
	scheme      [][]float32
}

type Records []Record

func (db *Records) Add(myScore, herScore float32, scheme [][]float32) error {
	neededScore := calculateScore(herScore, scheme)
	myPoints := (*db)[len(*db)-1].myPoints + 1
	herPoints := (*db)[len(*db)-1].herPoints + 1


	if myScore > neededScore {
		myPoints++
	} else {
		herPoints++
	}
	
	r := Record{
		weekID: len(*db),
		myScore: myScore,
		herScore: herScore,
		neededScore: neededScore,
		winForMe: myScore > neededScore,
		myPoints: myPoints,
		herPoints: herPoints,
		scheme: scheme,
	}
	*db = append(*db, r)
	return nil
}

func (db *Record) Get(fileName string) error {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	
	*db, err = parseTxt(f)
	if err != nil {
		return err
	}
	return nil
}

func (db *Record) Save(fileName string) error {
	return nil
}