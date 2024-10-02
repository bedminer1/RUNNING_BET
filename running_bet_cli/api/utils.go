package api

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func calculateScore(herScore float32, scheme [][]float32) float32 {
	var res float32 = 0

	for i := range scheme {
		dist := scheme[i][0]
		mult := scheme[i][1]
		var prevDist float32

		if i > 0 {
			prevDist = scheme[i-1][0]
		} else {
			prevDist = 0
		}
		span := dist - prevDist
		if herScore <= span {
			res += herScore * mult
			return res
		}

		res += span * mult
		herScore -= span
	}

	res += herScore
	return res
}

// incomplete because txt files previously didnt store scheme of a record
func parseTxt(data []byte) (Records, error) {
	res := Records{}
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		fields := strings.Split(l, " ")

		// parsing
		weekID, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		myScore, err := strconv.ParseFloat(fields[1], 32)
		if err != nil {
			return nil, err
		}
		herScore, err := strconv.ParseFloat(fields[2], 32)
		if err != nil {
			return nil, err
		}
		neededScore, err := strconv.ParseFloat(fields[3], 32)
		if err != nil {
			return nil, err
		}
		winForMe := myScore > neededScore
		myPoints, err := strconv.Atoi(fields[5])
		if err != nil {
			return nil, err
		}
		herPoints, err := strconv.Atoi(fields[6])
		if err != nil {
			return nil, err
		}

		r := Record{
			WeekID:      weekID,
			MyScore:     float32(myScore),
			HerScore:    float32(herScore),
			NeededScore: float32(neededScore),
			WinForMe:    winForMe,
			MyPoints:    myPoints,
			HerPoints:   herPoints,
		}
		res = append(res, r)
	}

	return res, nil
}

func parseJSON(data []byte) (Records, error) {
	res := Records{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// saving to SQLITE
const (
	createTable string = `CREATE TABLE records (
WeekID INTEGER PRIMARY KEY,
MyScore REAL,
HerScore REAL,
NeededScore REAL,
WinForMe BOOLEAN,
MyPoints INTEGER,
HerPoints INTEGER,
Scheme TEXT
);`
)

type dbRepo struct {
	db *sql.DB
	sync.RWMutex
}

func NewSQLite3Repo(dbfile string) (*dbRepo, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}

	return &dbRepo{
		db: db,
	}, nil
}

func (repo *dbRepo) Upsert(record Record) error {
	var exists bool
	err := repo.db.QueryRow("SELECT EXISTS(SELECT 1 FROM records WHERE weekID=?)", record.WeekID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err := repo.db.Exec(`
			UPDATE records SET my_score=?, her_score=?, needed_score=?, win_for_me=?, my_points=?, her_points=?, scheme=?
			WHERE weekID=?`,
			record.MyScore, record.HerScore, record.NeededScore, record.WinForMe, record.MyPoints, record.HerPoints, convertSchemeToBytes(record.Scheme), record.WeekID,
		)
		if err != nil {
			return err
		}
	} else {
		_, err := repo.db.Exec(`
			INSERT INTO records (weekID, my_score, her_score, needed_score, win_for_me, my_points, her_points, scheme) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			record.WeekID, record.MyScore, record.HerScore, record.NeededScore, record.WinForMe, record.MyPoints, record.HerPoints, convertSchemeToBytes(record.Scheme),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Helper function to convert Scheme to a byte array for storage in SQLite
func convertSchemeToBytes(scheme [][]float32) []byte {
	data, _ := json.Marshal(scheme)
	return data
}