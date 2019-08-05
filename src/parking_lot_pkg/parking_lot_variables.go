package parking_lot_pkg

type slot struct {
	vehicle_no string
	color      string
	//vehicle_type string //for extending the design to different types of vehicle
}

//type billing struct {

//}

var (
	parking_slot     = make(map[int]slot)
	park_number_slot = make(map[string]int)
	n                int
	slot_remaning    []int
)
