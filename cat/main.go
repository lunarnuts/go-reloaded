package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func Write(str string) {
	for _, p := range str {
		z01.PrintRune(p)
	}
}
func main() {
	s := os.Args
	s = s[1:]
	leng := 0
	for range s {
		leng++
	}
	if leng < 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return
		}
		/**go func(c []byte) {
			os.Stdout.Write(c)
			os.Exit(0)
		}(content)**/
		//os.Stdout.Write(content)
	}
	for _, n := range s {
		content, err := ioutil.ReadFile(n)
		if err != nil {
			j := "open " + string(n) + ": no such file or directory\n"
			Write(j)
			return
		}
		for _, p := range string(content) {
			z01.PrintRune(p)
		}
	}
}
