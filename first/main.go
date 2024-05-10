package main

import "fmt" // include

func main() {
	fmt.Println("Hello World!") // cout
	fmt.Println("Приветики")
	fmt.Println()

	var int_variablre8_min int8 = -128
	var int_variablre8_max int8 = 127
	var int_variablre16_min int16 = -32768               // 32.768
	var int_variablre16_max int16 = 32767                //  32.767
	var int_variablre32_min int32 = -2147483648          // 2.147.483.648
	var int_variablre32_max int32 = 2147483647           //  2.147.483.647
	var int_variablre64_min int64 = -9223372036854775808 // 9.223.372.036.854.775.808
	var int_variablre64_max int64 = 9223372036854775807  //  9.223.372.036.854.775.807
	var uint_variablre8_min uint8 = 0
	var uint_variablre8_max uint8 = 255
	var uint_variablre16_min uint16 = 0
	var uint_variablre16_max uint16 = 65535 // 65.535
	var uint_variablre32_min uint32 = 0
	var uint_variablre32_max uint32 = 4294967295 // 4.294.967.295
	var uint_variablre64_min uint64 = 0
	var uint_variablre64_max uint64 = 18446744073709551615 // 18.446.744.073.709.551.615

	fmt.Println("int_variablre8_min:  ", int_variablre8_min)
	fmt.Println("int_variablre8_max:   ", int_variablre8_max)
	fmt.Println("int_variablre16_min: ", int_variablre16_min)
	fmt.Println("int_variablre16_max:  ", int_variablre16_max)
	fmt.Println("int_variablre32_min: ", int_variablre32_min)
	fmt.Println("int_variablre32_max:  ", int_variablre32_max)
	fmt.Println("int_variablre64_min: ", int_variablre64_min)
	fmt.Println("int_variablre64_max:  ", int_variablre64_max)
	fmt.Println("uint_variablre8_min:  ", uint_variablre8_min)
	fmt.Println("uint_variablre8_max:  ", uint_variablre8_max)
	fmt.Println("uint_variablre16_min: ", uint_variablre16_min)
	fmt.Println("uint_variablre16_max: ", uint_variablre16_max)
	fmt.Println("uint_variablre32_min: ", uint_variablre32_min)
	fmt.Println("uint_variablre32_max: ", uint_variablre32_max)
	fmt.Println("uint_variablre64_min: ", uint_variablre64_min)
	fmt.Println("uint_variablre64_max: ", uint_variablre64_max)
	fmt.Println()

	var float_variable float64 = 123.456
	fmt.Println("float ", float_variable)
	fmt.Println()

	new_number := 136.458
	fmt.Println("test new init float", new_number)
	fmt.Println()

	var test_string string = "test 1"
	fmt.Println("test strings 1: ", test_string)
	test_string2 := "test test 2"
	fmt.Println("test strings 2: ", test_string2)
	fmt.Println("size1: ", len(test_string), " size2: ", len(test_string2))
	fmt.Println()

	fmt.Println("Test input string")
	var test_strig_3 string
	fmt.Println("Please enter string")
	fmt.Scan(&test_strig_3)
	// fmt.Sprint(number) convert number to string
	fmt.Println("your string: '", test_strig_3, "', it's size is: ", len(test_strig_3))
	fmt.Println("Please enter another string")
	fmt.Scanln(&test_strig_3)
	fmt.Println("your next string: <", test_strig_3, ">, it's size is: ", len(test_strig_3))
	fmt.Println()

	var t1 int8 = 13
	var t2 int64 = int64(t1)
	fmt.Println(t2)
}
