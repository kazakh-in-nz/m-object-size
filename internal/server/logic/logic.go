package logic

var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

func GetSize() float64 {
	oldScore := pastFourScores[0] + pastFourScores[1]
	newScore := pastFourScores[2] + pastFourScores[3]

	diff := newScore - oldScore

	if diff > 0 {
		size := 600.0 + diff*60

		if size < 2000.0 {
			return size
		}

		return 2000.0
	}

	if diff > -5.0 && diff <= 0.0 {
		return 100.0 + 18*diff
	}

	return 10.0
}

func SetScore(a float64) bool {
	pastFourScores = append(pastFourScores, a)
	pastFourScores = pastFourScores[1:]
	return true
}
