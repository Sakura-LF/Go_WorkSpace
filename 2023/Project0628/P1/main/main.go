package main

import "fmt"

func main() {
	arr := [6]int{1, 7, 3, 6, 5, 6}

	fmt.Println(Test1(arr))

}

func Test1(nums [6]int) int {

	for key, _ := range nums {
		//求左和
		leftsum := 0
		for i := key; i > 0; i-- {
			if i == 0 {
				leftsum = 0
			} else {
				leftsum += nums[i-1]
			}
		}

		//求右和
		rightsum := 0
		for i := key; i < len(nums)-1; i++ {
			if i == len(nums)-1 {
				rightsum = 0
			} else {
				rightsum += nums[i+1]
			}
		}

		if rightsum == leftsum {
			return key
		}
	}
	return -1
}

func Test(nums []int) int {
	//sum := 0
	//for _, v := range nums {
	//	sum += v
	//}
	//leftsum := 0 //左侧数之和
	//for key, _ := range nums {
	//	sum -= nums[key] //右侧数之和
	//
	//	if leftsum == sum { //判断左侧数字是否等于右侧数
	//		return key
	//	}
	//	//在判断之后计算左侧数:可以就可以不包含本索引
	//	leftsum += nums[key]
	//}
	//return -1
	//leftsum := 0
	//rightsum := 0
	//for key, value := range nums {
	//	//求左和
	//	for i := key; i >= 0; i-- {
	//		leftsum += nums[i]
	//	}
	//
	//	//求右和
	//}

	return -1
}
