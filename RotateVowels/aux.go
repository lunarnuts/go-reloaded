package student

func CountWords(str string, charset string) int {
	l := Length(str)
	l2 := Length(charset)
	count := 0
	p := 0
	for n := range str {
		if n != 0 && n != l {
			j := ""
			if n+l2 <= l {
				j = string(str[n : n+l2])
			}
			if j == charset {
				str2 := string(str[p:n])
				if str2 != "" {
					count++
				}
				p = n + l2
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" {
		count++
	}
	return count
}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	power--
	return nb * RecursivePower(nb, power)
}

func Index(s string, c rune) int {
	str := []rune(s)
	index := -1
	for n := range str {
		if str[n] == c {
			index = n
			return index
		}
	}
	return index
}

func Valid(s string) bool {
	if Length(s) < 2 {
		return false
	}
	if (Index(s, '-') != -1) || (Index(s, '+') != -1) {
		return false
	}
	count := 0
	for _, a := range s {
		count++
		if Index(s[count:], a) >= 0 {
			return false
		}
	}
	return true
}

func Length(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func IntLen(str []int) int {
	s := 0
	for range str {
		s++
	}
	return s
}

func MyAppend(a, b []int) []int {
	p := [200]int{-1}
	count := 0
	for range a {
		p[count] = a[count]
		count++
	}
	len2 := IntLen(a)
	for range b {
		p[count] = b[count-len2]
		count++
	}
	return p[:count]
}
func MyAppendRune(a, b []rune) []rune {
	p := [200]rune{-1}
	count := 0
	for range a {
		p[count] = a[count]
		count++
	}
	len2 := 0
	for range a {
		len2++
	}
	for range b {
		p[count] = b[count-len2]
		count++
	}
	return p[:count]
}
