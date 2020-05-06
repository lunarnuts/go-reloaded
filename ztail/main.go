package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	args := os.Args[1:]
	length := 0
	for range args {
		length++
	}
	if length < 1 {
		/**var data []byte
		for {
			_, err := os.Stdin.Read(data)
			if err != nil {
				continue
			} else {
				os.Stdout.Write(data)
				defer os.Exit(0)
				break

			}
		}
		**/
		p, _ := student.MyRead()
		os.Stdout.Write(p)
		os.Exit(0)
	}
	num := 0
	flagc := false //flag to know if -c is used
	flagp := false //flag to know if after -c option user put positive sign
	//or negative before int
	for i, arg := range args {
		if arg == "-c" {
			if student.CheckS(args[i+1]) {
				if args[i+1][0] == '+' {
					flagp = true
				}
				num = student.Atoi(args[i+1])
				length = length - 1
				flagc = true
			} else {
				fmt.Printf("tail: invalid number of bytes: '%s'\n", args[i+1])
				os.Exit(1)
			}
		} else if flagc {
			length = length - 1
			flagc = false
		}
	}
	flag2 := false
	errflag := false
	cl := 0
	message := ""
	for j, fname := range args { //open files in args
		if fname == "-c" {
			flag2 = true
			continue
		}
		if flag2 && args[j-1] == "-c" {
			continue
		}
		flag2 = true
		file, err := os.Open(fname)

		if err != nil {
			message += "tail: cannot open '" + fname + "' for reading: No such file or directory\n"
			errflag = true
			continue
		}
		if cl != 0 {
			message += "\n"
		}
		if length > 1 {
			message += "==> " + fname + " <==\n"
			cl++
		}

		fileinfo, err := file.Stat()
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			errflag = true
			return
		}

		fs := fileinfo.Size()
		buffer := student.MyMake(int(fs))

		lenf, err := file.Read(buffer)
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			errflag = true
			return
		}
		if num == 0 || num > lenf {
			count := 0
			ind := 0
			for i := lenf - 1; i >= 0; i-- {
				if buffer[i] == 10 {
					count++
				}
				if count == 10 {
					ind = i
					break
				}
			}
			message += string(buffer[ind:])

		} else {
			if flagp {
				message += string(buffer[num-1:])
			} else {
				message += string(buffer[lenf-num:])
			}
		}
	}
	fmt.Printf(message)
	if errflag {
		os.Exit(1)
	}

	os.Exit(0)

}
