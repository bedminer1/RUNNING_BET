package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
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
	createTable string = `CREATE TABLE IF NOT EXISTS records (
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
			UPDATE records SET MyScore=ROUND(?, 2), HerScore=ROUND(?, 2), NeededScore=ROUND(?, 2), WinForMe=?, MyPoints=?, HerPoints=?, Scheme=?
			WHERE WeekID=?`,
			record.MyScore, record.HerScore, record.NeededScore, record.WinForMe, record.MyPoints, record.HerPoints, convertSchemeToBytes(record.Scheme), record.WeekID,
		)
		if err != nil {
			return err
		}
	} else {
		_, err := repo.db.Exec(`
			INSERT INTO records (WeekID, MyScore, HerScore, NeededScore, WinForMe, MyPoints, HerPoints, Scheme) 
			VALUES (?, ROUND(?, 2), ROUND(?, 2), ROUND(?, 2), ?, ?, ?, ?)`,
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

func (repo *dbRepo) Fetch() (Records, error) {
	rows, err := repo.db.Query("SELECT * FROM records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res Records
	for rows.Next() {
		var (
			record     Record
			schemeText string
		)
		err := rows.Scan(&record.WeekID, &record.MyScore, &record.HerScore, &record.NeededScore, &record.WinForMe, &record.MyPoints, &record.HerPoints, &schemeText)
		if err != nil {
			return nil, err
		}

		// deserialize scheme from string to [][]float32
		err = json.Unmarshal([]byte(schemeText), &record.Scheme)
		if err != nil {
			return nil, err
		}

		res = append(res, record)
	}

	return res, nil
}

func printRecords(out io.Writer, records Records) {
	// Create a new tabwriter to format the table output
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', tabwriter.Debug)

	// Print table headers
	fmt.Fprintf(w, "WeekID\tMyScore\tHerScore\tNeededScore\tWinForMe\tMyPoints\tHerPoints\tScheme\n")
	fmt.Fprintf(w, "------\t-------\t--------\t-----------\t---------\t---------\t---------\t------\n")

	// Iterate through the slice of records
	for _, record := range records {
		// Print each record in the table
		fmt.Fprintf(w, "%d\t%.2f\t%.2f\t%.2f\t%t\t%d\t%d\t%v\n", record.WeekID, record.MyScore, record.HerScore, record.NeededScore, record.WinForMe, record.MyPoints, record.HerPoints, record.Scheme)
	}

	// Flush the writer to ensure all rows are printed
	w.Flush()
}
