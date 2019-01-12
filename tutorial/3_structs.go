package main

// USIXTEENBITMAX must have comment
const USIXTEENBITMAX float64 = 65535

// KMHMULITIPLE must have comment
const KMHMULITIPLE float64 = 1.60934

// Car test struct
type Car struct {
	gasPedal      uint16
	brakePadel    uint16
	steeringWhell int16
	topSpeedKMH   float64
}

// Kmh z这是car的成员方法，是在外部通过关联的, what is value receiver
func (c Car) Kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / USIXTEENBITMAX)
}

// Mmh z这是car的成员方法，是在外部通过关联的
func (c Car) Mmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / USIXTEENBITMAX / KMHMULITIPLE)
}

// NewTopSpeed pointer receiver
func (c *Car) NewTopSpeed(newSpeed float64) {
	c.topSpeedKMH = newSpeed
}

// func main() {
// 	ACar := Car{gasPedal: 65000,
// 		brakePadel:    0,
// 		steeringWhell: 1234,
// 		topSpeedKMH:   231.9}

// 	fmt.Println("ACar's gasPedal:", ACar.gasPedal)
// 	fmt.Println("ACar's Kmh:", ACar.Kmh())
// 	fmt.Println("ACar's mmh:", ACar.Mmh())

// 	ACar.NewTopSpeed(500)
// 	fmt.Println("ACar's Kmh:", ACar.Kmh())
// 	fmt.Println("ACar's mmh:", ACar.Mmh())
// }
