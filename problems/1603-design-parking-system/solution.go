type ParkingSystem struct {
	Slot []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	p := ParkingSystem{}
	p.Slot = []int{big, medium, small}
	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Slot[carType-1] == 0 {
		return false
	}

	this.Slot[carType-1] -= 1
	return true
}

