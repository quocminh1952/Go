package main

import (
	"fmt"
)

type vehicle struct {
	name   string
	action string
}

// tạo method liên kết với với kiểu vehicle
// car là reciver(tham số đặc biệt) đại diện cho instance của vehicle
func (car vehicle) run() {
	fmt.Println("run" + car.name)
}

func (plane vehicle) fly() {
	fmt.Println("fly" + plane.name)
}

func main() {
	VinFast := vehicle{
		name:   "VF5",
		action: "run Electric",
	}

	VietNamAir := vehicle{
		name:   "Bo024",
		action: "run Electric",
	}

	VinFast.fly()
	VinFast.run()
	VietNamAir.fly()
}
