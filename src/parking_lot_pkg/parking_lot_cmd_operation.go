package parking_lot_pkg

import (
	"bufio"
	"fmt"
	"os"
)

func CMD_Operation_fn() {
	var l_err error
	var input_cmd string
	for {
		reader := bufio.NewReader(os.Stdin)
		input_cmd, l_err = reader.ReadString('\n')
		if l_err != nil {

			fmt.Println(l_err)
			return

		}
		l_err = Process_fn(input_cmd)
		if l_err != nil {
			return

		}
	}

}
