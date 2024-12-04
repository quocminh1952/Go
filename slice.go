package main

import (
	"fmt"
)

// Slice : một phần căắt từ mảng
//
//	có thể chia sẻ dữ liệu với mảng gốc
//	là một mảng động
func main() {
	var arr = [5]string{"ă", "b", "c", "d", "e"}

	// I : Khai báo slice
	//1: Tạo slice
	slice := arr[1:4] // từ phần tử thứ 1 -> 3 ( mảng bắt đầu từ 0 )
	// slice và arr dùng chung một địa chỉ nhớ
	fmt.Println(slice) // [b c d]

	//2: Khai báo trực tiếp
	slice2 := []string{"Minh", "Nguyen"}
	fmt.Println(slice2)

	//3: Sử dụng make
	slice3 := make([]int, 3, 5) // length:3  capacity:5
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//II: Thao tác cơ bản với slice
	//1.Thêm Phần tử vào slice
	slice = append(slice, "f", "g")
	fmt.Println(slice)
	// slice sử dụng append sẽ thay đổi ô nhớ <=> không chung ô nhớ với arr

	//2.Truy cập thay đổi giá trị
	fmt.Println(slice[0])
	slice[0] = "X"
	fmt.Println(slice)
	fmt.Println(arr) // khác với slice vì slice đã sử dụng append

	//3.Duyệt qua slice
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// slice default
	// arr = ["ă", "b", "c", "d", "e"]
	slice4 := arr[:4]
	fmt.Println(slice4)

	slice5 := arr[2:]
	fmt.Println(slice5)

}
