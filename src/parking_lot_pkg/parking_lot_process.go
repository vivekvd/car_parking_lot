package parking_lot_pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Process_fn(input_cmd string) (r_err error) {
	input_cmd = strings.TrimSuffix(input_cmd, "\n")
	cmd_arr := strings.Split(input_cmd, " ")

	switch cmd_arr[0] {
	case "create_parking_lot":

		n, r_err = strconv.Atoi(cmd_arr[1])

		if r_err != nil {
			fmt.Println(r_err)
			r_err = errors.New("Unable to Convert")
			return
		}

		slot_remaning = make([]int, n)
		for i := 0; i < n; i++ {
			slot_remaning[i] = i + 1
		}
		fmt.Println("Created a parking lot with", n, " slots")
		fmt.Println(len(slot_remaning))
	case "park":

		if len(slot_remaning) > 0 {
			var vehicle_check bool
			slot_no := slot_remaning[0]
			temp := slot{vehicle_no: cmd_arr[1], color: cmd_arr[2]}
			for index := range parking_slot {
				slot_details := parking_slot[index]
				if slot_details.vehicle_no == cmd_arr[1] {
					vehicle_check = true
					break
				}
			}
			if !vehicle_check {
				slot_remaning = slot_remaning[:0+copy(slot_remaning[0:], slot_remaning[0+1:])]
				parking_slot[slot_no] = temp
				park_number_slot[cmd_arr[1]] = slot_no
				// Capture entry in the transaction audit_logs
				fmt.Println("Allocated slot number: ", slot_no)
			} else {
				fmt.Println("Vehicle no already parked")
			}

		} else {
			fmt.Println("Sorry, parking lot is full")
		}
	case "leave":
		var slot_no int
		slot_no, r_err = strconv.Atoi(cmd_arr[1])
		if r_err != nil {
			fmt.Println(r_err)
			r_err = errors.New("Unable to Convert")
			return
		}
		// Capture entry in the transaction audit_logs
		// Process_bill_fn()
		temp_slot := parking_slot[slot_no]
		delete(park_number_slot, temp_slot.vehicle_no)
		delete(parking_slot, slot_no)

		fmt.Println("Slot no ", slot_no, "  free")
		slot_remaning = append(slot_remaning, slot_no)
	case "status":

		fmt.Println("Slot No   Registration No   Colour")
		for slot_no, slot_info := range parking_slot {
			fmt.Println(slot_no, "     ", slot_info.vehicle_no, "      ", slot_info.color)
		}
	case "slot_number_for_registration_number":
		l_slot_number, ok := park_number_slot[cmd_arr[1]]

		if !ok {
			fmt.Println("Not found")
		} else {
			fmt.Println(l_slot_number)
		}
	case "registration_numbers_for_cars_with_colour":
		l_flag := false
		for _, l_slot_info := range parking_slot {
			if strings.EqualFold(cmd_arr[1], l_slot_info.color) {
				l_flag = true
				fmt.Print(l_slot_info.vehicle_no, "  ")
			}
		}
		if !l_flag {
			fmt.Println("Not Found")
		} else {
			fmt.Println("")
		}
	case "slot_numbers_for_cars_with_colour":
		l_flag := false
		for l_slot_number, l_slot_info := range parking_slot {

			if strings.EqualFold(l_slot_info.color, cmd_arr[1]) {
				l_flag = true
				fmt.Print(l_slot_number, "  ")
			}
		}
		if !l_flag {
			fmt.Println("Not Found")
		} else {
			fmt.Println("")
		}
	}

	return
}
