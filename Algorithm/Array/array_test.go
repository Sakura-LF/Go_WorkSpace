package Array

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{-1, 0, 3, 9, 11, 13, 22, 27, 33, 57, 66, 77}
	fmt.Println(BinarySearch2(arr, 13))
}

func TestRemoveElement(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	fmt.Println(RemoveElement(arr, 3))
}

func TestRemoveElement2(t *testing.T) {
	arr := []int{1, 3, 2, 6, 2, 3, 5, 3}
	fmt.Println(RemoveElement2(arr, 3))
}
func TestSquare(t *testing.T) {
	arr := []int{1, 3, 2, 6, 2, 3, 5, 3}
	Square(arr)
}
