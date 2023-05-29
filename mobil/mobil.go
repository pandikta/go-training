package mobil

type Car struct {
	Tires [4]Tires
	Doors [2]Doors
}

type Tires interface {
	CanRolling() bool
}

func (c Car) ChangeTire(i int, tire Tires) {
	if i >= 4 {
		return
	}

	c.Tires[i] = tire
}
