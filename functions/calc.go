package linearstats

import (
	"math"
)

func Average(data []int) float64 {
	var somme float64
	for i := 0; i < len(data); i++ {
		somme += float64(data[i])
	}

	return somme / float64(len(data))
}

func Median(data []int) float64 {
	return float64(len(data)-1) / 2
}

func LinearRegression(data []int) (float64, float64) {
	if len(data) == 0 {
		return 0.0, 0.0
	}

	average := Average(data)
	median := Median(data)

	var CovarianceXY float64
	var VarianceX float64
	for i := 0; i < len(data); i++ {
		CovarianceXY += (float64(data[i]) - average) * (float64(i) - median)
		VarianceX += math.Pow(float64(i)-median, 2)
	}

	a := CovarianceXY / VarianceX
	b := average - (a * median)

	return a, b
}

func PearsonCorrelationCoefficient(data []int) float64 {
	if len(data) == 0 {
		return 0.0
	}

	var sumXY float64
	var sumX float64
	var sumY float64
	var sumPowX float64
	var sumPowY float64

	length := len(data)
	for i := 0; i < length; i++ {
		sumXY += float64(data[i]) * float64(i)
		sumX += float64(i)
		sumY += float64(data[i])
		sumPowX += math.Pow(float64(i), 2)
		sumPowY += math.Pow(float64(data[i]), 2)
	}

	a := (float64(length) * (sumXY)) - (sumX * sumY)
	b := (float64(length)*sumPowX - math.Pow(sumX, 2)) * (float64(length)*sumPowY - math.Pow(sumY, 2))

	return a / math.Sqrt(b)
}
