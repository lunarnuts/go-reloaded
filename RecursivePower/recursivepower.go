package student

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	power--
	return nb * RecursivePower(nb, power)
}
