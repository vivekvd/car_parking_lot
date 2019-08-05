package parking_lot_pkg

func Start_fn(file_name string) {
	if file_name != "" {
		File_Operation_fn(file_name)

	} else {
		CMD_Operation_fn()
	}
	return
}
