package student

func AtoiBase(s string, base string) int {
	l := 1
	length := Length(s)
	if length < 1 {
		return 0
	}
	if s[0] == '-' {
		l = -1
		s = s[1:]
	} else if s[0] == '+' {
		l = 1
		s = s[1:]
	}
	length = Length(s)
	if length < 1 {
		return 0
	}
	valid := Valid(base)
	sum := 0
	if valid == false {
		return 0
	}
	sequence := []int{}
	var r int
	count := 0
	for _, n := range s {
		r = Index(base, n)
		if r < 0 {
			return 0
		}
		sequence = MyAppend([]int{r}, sequence)
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
