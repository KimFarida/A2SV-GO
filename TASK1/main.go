package main
import "fmt"

/*
** Fundamentals of Go Tasks
Task: Sum of Numbers
Write a Go function that takes a slice of integers as input and returns the sum of all the numbers. If the slice is empty, the function should return 0.
[Optional]: Write a test for your function

*/

func sumNum(nums[]int)int{

	total := 0
	for _, num := range nums{
		total+= num
	}
	return total
}

func main(){
	test_1 := []int{
		1, 2, 3, 4, 5,
	}
	test_2 := make([]int, 0)
	fmt.Println(sumNum(test_1))
	fmt.Println(sumNum(test_2))
}