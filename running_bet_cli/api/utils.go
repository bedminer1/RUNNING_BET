package api

import (
	"strconv"
	"strings"
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
