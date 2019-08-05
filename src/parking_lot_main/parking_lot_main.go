package main

import (
	"os"
	"parking_lot_pkg"
)

func main() {
	var file_name string
	args := os.Args
	if len(args) > 1 {
		file_name = os.Args[1]
	}
	parking_lot_pkg.Start_fn(file_name)
	return
}
