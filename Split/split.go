package piscine

func length(str string) int {
	l := 0
	for range str {
		l++
	}
	return l
}
func CountWords(str string, charset string) int {
	l := length(str)
	l2 := length(charset)
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
func Split(str string, charset string) []string {
	s := make([]string, CountWords(str, charset))
	l := length(str)
	count := 0
	l2 := length(charset)
	p := 0
	for n := range str {
		if /*(n != 0 &&*/ n != l {
			j := ""
			if n+l2 <= l {
				j = string(str[n : n+l2])
			}
			if j == charset {
				str2 := string(str[p:n])
				if str2 != "" && str2 != charset {
					s[count] = str2
					count++
				}
				p = n + l2
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" && str2 != charset {
		s[count] = str2
	}
	return s
}
