package common

func CalFactorial(n float64) float64 {
	r := float64(1)

	for n > 1 {
		// fmt.Println(n,r)

		r *= n
		n -= 1
		// fmt.Println(n,r)

	}
	
	return r
}