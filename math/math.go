package math

func Lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*t
}
