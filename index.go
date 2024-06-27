package main

import "fmt"

func main() {
	//Khai báo biến bằng var
	var boolVar bool = true

	//Khai không khai báo kiểu dữ liệu thì complier cũng có thể hiểu và tự động gán dữ liệu cho biến
	var stringVar = "Hello world"
	var string2Var string // Value của string sẽ là ""

	// Khai báo biến kiểu rút gọn
	// Giá trị nguyên sẽ tự động được gán kiểu int
	intVar := 1 // Kiểu int sẽ tự động được gán size dựa vào target machine lúc compile. Thông thường sẽ là int64
	//Hoặc có thể khai báo nhiều biến cùng lúc
	var (
		int8Var  int8  = 1
		int16Var int16 = 3419
		int32Var int32 = 1823904
		int64Var int64 = 123091247678678
	)

	//Khi assign 1 biến int với giá trị của int64, compiler sẽ báo lỗi mặc dù int là 64bit
	intVar = int64Var // cannot use int64Var (variable of type int64) as int value in assignment
	// Để assign 1 biến int với giá trị của biến int64 thì cần phải convert
	intVar = (int(int64Var))

	var uintVar uint = 12351234123948120934 //Tương tự như int, unit64
	// Khi phép gán vượt giá trị, compiler sẽ báo lỗi
	var uint8Var uint8 = 256
	var uint16Var uint16 = 1823
	var uint32Var uint32 = 123124134
	var uint64Var uint64 = 1234871234

	//Hằng số
	const uintPtrVar uintptr = 1234812934092134092 // Đây là kiểu số nguyên dương đủ to để chứa giá trị của trỏ. phụ thuộc vào target machine uint64

	var float32Var float32 = 1.2
	//Khi dùng shorthand thì float sẽ là float64
	floatVar := 1.2 // float64

	var complex64Var complex64 = -4 + 12i
	// Khi dùng shorthand thì complex sẽ tự động gán complex128
	complexVar := 123 + 1243i //complex128

	var p *int
	i := 542
	p = &i
	fmt.Println(*p) // 452
	*p = 21         // i cũng bằng 21
	fmt.Println(*p) // 21

	fmt.Println(boolVar)
	fmt.Println(stringVar)
	fmt.Println(string2Var)
	fmt.Println(intVar)
	fmt.Println(int8Var)
	fmt.Println(int16Var)
	fmt.Println(int32Var)
	fmt.Println(uintVar)
	fmt.Println(uint8Var)
	fmt.Println(uint16Var)
	fmt.Println(uint32Var)
	fmt.Println(uint64Var)
	fmt.Println(uintPtrVar)
	fmt.Println(floatVar)
	fmt.Println(float32Var)
	fmt.Println(complex64Var)
	fmt.Println(complexVar)
}
