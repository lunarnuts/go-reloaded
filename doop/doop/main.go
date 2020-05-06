package main

import (
	"fmt"
	"os"

	student ".."
)

var INT_MAX int64 = 9223372036854775807
var INT_MIN int64 = -9223372036854775808

func Sum(si_a, si_b int64) int64 {

	if ((si_b > 0) && (si_a > (INT_MAX - si_b))) ||
		((si_b < 0) && (si_a < (INT_MIN - si_b))) {
		return 0
	}
	return si_a + si_b
}
func Diff(si_a, si_b int64) int64 {
	if (si_b > 0 && si_a < INT_MIN+si_b) ||
		(si_b < 0 && si_a > INT_MAX+si_b) {
		return 0
	}
	return si_a - si_b
}
func Multi(si_a, si_b int64) int64 {
	if si_a > 0 { /* si_a is positive */
		if si_b > 0 { /* si_a and si_b are positive */
			if si_a > (INT_MAX / si_b) {
				return 0
			}
		} else { /* si_a positive, si_b nonpositive */
			if si_b < (INT_MIN / si_a) {
				return 0
			}
		} /* si_a positive, si_b nonpositive */
	} else { /* si_a is nonpositive */
		if si_b > 0 { /* si_a is nonpositive, si_b is positive */
			if si_a < (INT_MIN / si_b) {
				return 0
			}
		} else { /* si_a and si_b are nonpositive */
			if (si_a != 0) && (si_b < (INT_MAX / si_a)) {
				return 0
			}
		} /* End if si_a and si_b are nonpositive */
	} /* End if si_a is nonpositive */

	return si_a * si_b
}
func Div(si_a, si_b int64) int64 {
	if si_b == 0 {
		fmt.Printf("No division by 0\n")
		os.Exit(1)
	}
	if (si_a == INT_MIN) && (si_b == -1) {
		return 0
	}
	return si_a / si_b
}
func Mod(si_a, si_b int64) int64 {
	if si_b == 0 {
		fmt.Printf("No remainder of division by 0\n")
		os.Exit(1)
	}
	if (si_a == INT_MIN) && (si_b == -1) {
		return 0
	}
	return si_a % si_b
}
func Check(s string) bool {
	Che := true
	num := 0
	for _, n := range s {
		if (n < '0') || (n > '9') {
			num++
		}
	}

	if num > 0 {
		Che = false
	}
	return Che
}
func main() {
	s := os.Args[1:]
	l := 0
	for range s {
		l++
	}
	var result int64 = 0
	if l == 3 {
		p := s[0]
		q := s[2]
		if p[0] == '-' {
			p = p[1:]
		}
		if q[0] == '-' {
			q = q[1:]
		}
		if Check(p) && Check(q) {
			n1, ov1 := student.Atoi(s[0])
			n2, ov2 := student.Atoi(s[2])
			if ov1 == true || ov2 == true {
				fmt.Printf("0\n")
				os.Exit(1)
			}
			switch s[1] {
			case "/":
				result = Div(int64(n1), int64(n2))
			case "*":
				result = Multi(int64(n1), int64(n2))
			case "+":
				result = Sum(int64(n1), int64(n2))
			case "-":
				result = Diff(int64(n1), int64(n2))
			case "%":
				result = Mod(int64(n1), int64(n2))
			default:
				fmt.Printf("0\n")
				os.Exit(1)
			}
		}
	} else {
		os.Exit(0)
	}
	fmt.Printf("%d\n", result)
	os.Exit(0)
}
