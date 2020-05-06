package student

func AtoiBase(s string, base string) int {
	var l int = 1
	switch s[0] {
	case '-':
		l = -1
		s = s[1:]
	case '+':
		l = 1
		s = s[1:]
	}
	if Length(s) < 1 {
		return 0
	}
	sum := 0
	if !Valid(base) {
		return 0
	}
	sequence := []int{}
	count := 0
	for _, n := range s {
		if Index(base, n) < 0 {
			return 0
		}
		sequence = MyAppend([]int{Index(base, n)}, sequence)
		count++
	}
	for a := range sequence {
		sum = sum + sequence[a]*RecursivePower(Length(base), a)
	}
	if Index(base, '-') >= 0 || Index(base, '+') >= 0 {
		sum = 0
	}
	return sum * l
}
