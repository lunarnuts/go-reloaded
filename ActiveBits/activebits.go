package student

func ActiveBits(n int) uint {
	var count uint = 0
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return count
}
