package api

import "fmt"

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

func parseTxt(data []byte) (Record, error) {
	fmt.Println(data)
	return Record{}, nil
}
