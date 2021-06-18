package main

import (
	"fmt"
)

func main() {
	// Variable
	// Variable can be declare as:
	// var varibale_name datatype = value
	// or varible_name := value
	var s string = "This is go lang string"
	var i int = 10
	var f float32 = 100.45
	var d float64 = 10003.3332
	var b bool = true
	fmt.Printf("string: %v\n int: %v\n float32: %v\n float64: %v\n bool: %v\n\n", s, i, f, d, b)
	// or
	ss := "This is go lang string"
	ii := 10
	ff := 100.45
	dd := 10003.3332
	bb := false
	fmt.Printf("string: %v\n int: %v\n float32: %v\n float64: %v\n bool: %v\n\n", ss, ii, ff, dd, bb)

	// Constant
	// Declare constant as:
	// const constant_name datatype = value

	const sc string = "This is go lang string"
	const ic int = 10
	const fc float32 = 100.45
	const dc float64 = 10003.3332
	const bc bool = true
	fmt.Printf("string: %v\n int: %v\n float32: %v\n float64: %v\n bool: %v\n\n", sc, ic, fc, dc, bc)

	// datatype in golang
	var p string = "This is strig datatype"
	var q int = 10
	var qq int8 = 1
	var qqq int16 = 111
	var qqqq int32 = 11213
	var qqqqq int64 = 12343939
	var r float32 = 344.55
	var rr float64 = 44343.999993
	fmt.Println(p, q, qq, qqq, qqqq, qqqqq, r, rr)

}
