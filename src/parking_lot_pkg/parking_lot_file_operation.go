package parking_lot_pkg

import (
	"bufio"
	"fmt"
	"os"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func File_Operation_fn(file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	input_cmd, e := Readln(r)
	for e == nil {
		l_err := Process_fn(input_cmd)
		if l_err != nil {
			fmt.Println("Error ", l_err)
			return
		}
		input_cmd, e = Readln(r)
	}
	return
}
