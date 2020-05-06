package student

func Split(str string, charset string) []string {
	s, l := make([]string, CountWords(str, charset)), Length(str)
	count, p, l2 := 0, 0, Length(charset)
	for n := range str {
		if n != 0 && n != l {
			j := ""
			if n+l2 <= l {
				j = string(str[n : n+l2])
			}
			if j == charset {
				str2 := string(str[p:n])
				if str2 != "" {
					s[count] = str2
					count++
				}
				p = n + l2
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" {
		s[count] = str2
	}
	return s
}
