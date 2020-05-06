package student

func CountWordsSpaces(str string) int {
	l := Length(str)
	count, p := 0, 0
	for n := range str {
		if n != 0 && n != l {
			j := string(str[n])
			if j == " " || j == "\t" || j == "\n" {
				str2 := string(str[p:n])
				if str2 != "" {
					count++
				}
				p = n + 1
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" {
		count++
	}
	return count
}
func SplitWhiteSpaces(str string) []string {
	s, l := make([]string, CountWordsSpaces(str)), Length(str)
	count, p := 0, 0
	for n := range str {
		if n != 0 && n != l {
			j := string(str[n])
			if j == " " || j == "\t" || j == "\n" {
				str2 := string(str[p:n])
				if str2 != "" && str2 != " " && str2 != "\t" && str2 != "\n" {
					s[count] = str2
					count++
				}
				p = n + 1
			}
		}
	}
	str2 := string(str[p:l])
	if str2 != "" && str2 != " " && str2 != "\t" && str2 != "\n" {
		s[count] = str2
	}
	return s
}
