package main

import (
    "fmt"
    "time"
    "math/rand"
)

var SortArray = []int{} // 待排序的数组
// 递归快排
func QuickSort(SortArray []int, low, high int) {
	if low < high {
		mid := Partition(SortArray, low, high) // 中间位置
		QuickSort(SortArray, low, mid - 1)
		QuickSort(SortArray, mid + 1, high)
	}
}
// 数组分区
func Partition(SortArray []int, low, high int) int {
	pivot := SortArray[high] // 基准值
	i := low - 1
	for j := low; j < high; j++ {
		if SortArray[j] <= pivot {
			i++
			SortArray[i], SortArray[j] = SortArray[j], SortArray[i] // 交换i，j两位置的值
		}
	}
    i++
	SortArray[i], SortArray[high] = SortArray[high], SortArray[i]
	return i
}

func main() {
    fmt.Printf("\nBefore sort:\n")
    for i :=0; i < 10; i++ { // 随机生成不同的10个数字数组
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        SortArray = append(SortArray, r.Intn(100))
        fmt.Printf("%d ", SortArray[i])
    }
	QuickSort(SortArray, 0, 9)
    fmt.Printf("\nAfter sort:\n")
    for _, value := range SortArray {
		fmt.Printf("%d ", value)
	}
}