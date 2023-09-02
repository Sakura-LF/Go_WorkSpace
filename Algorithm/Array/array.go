package Array

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, target int) int {
	//第一种写法,左闭右闭,[left,right]
	left := 0
	right := len(arr) - 1

	for left <= right {
		//防止溢出,等价于 (left + right) / 2
		middle := left + (right-left)/2

		if arr[middle] < target {
			//左区间没有要找的值
			left = middle + 1
		} else if arr[middle] > target {
			//右区间没有要找的值
			right = middle - 1
		} else {
			//中间值不大于也不小于target,则相等,返回索引
			return middle
		}
	}
	return -1
}

func BinarySearch2(arr []int, target int) int {
	//第二种写法,左闭右开[left,right)
	left := 0
	right := len(arr)

	for left < right {
		middle := left + (right-left)/2
		if arr[middle] < target {
			left = middle + 1
		} else if arr[middle] > target {
			right = middle
		} else {
			return middle
		}
	}
	return -1
}

func RemoveElement(arr []int, values int) int {
	lenght := len(arr)
	for i := 0; i < lenght; i++ {

		if values == arr[i] {
			// 将所有元素向前移动一位
			for j := i + 1; j < lenght; j++ {
				arr[j-1] = arr[j]
			}
			// 每个元素向前移动后,原本是i的元素也会向前
			i--
			// 因为移除(覆盖)了一个元素,所以length-1
			lenght--
		}
	}
	return lenght
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func RemoveElement2(arr []int, value int) int {
	// 定义慢指针
	slowIndex := 0
	for FastIndex := 0; FastIndex < len(arr); FastIndex++ {
		if arr[FastIndex] != value {
			// 如
			arr[slowIndex] = arr[FastIndex]
			slowIndex++
		}
	}
	arr = arr[:slowIndex]
	fmt.Println(arr)
	return slowIndex

	//length := len(arr)
	//res := 0
	//for i := 0; i < length; i++ {
	//	if arr[i] != value {
	//		arr[res] = arr[i]
	//		res++
	//	}
	//}
	//fmt.Println(arr)
	////切片,从0开始切到5,不包含5
	//arr=arr[:res]
	//return res

}

// 时间复杂度 O(n + nlogn)
func Square(arr []int) {
	// 遍历每个元素,并平方
	for key, value := range arr {
		arr[key] = value * value
	}
	//排序,go比较器
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
