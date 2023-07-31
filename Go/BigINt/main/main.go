package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 第一步：创建一个超大整型的一个对象
	var v1 big.Int
	var v2 big.Int

	// 第二步：在超大整型对象中写入值
	v1.SetInt64(9223372036854775807)
	fmt.Println(v1)

	v2.SetString("92233720368547758089223372036854775808", 10)
	fmt.Println(v2)

	// 第一步：创建一个超大整型的一个对象
	v3 := new(big.Int)
	v4 := new(big.Int)

	// 第二步：在超大整型对象中写入值
	v3.SetInt64(9223372036854775807)
	fmt.Println(v3)

	v4.SetString("92233720368547758089223372036854775808", 10)
	fmt.Println(v4)
}
