package slicetest

import "fmt"

func Example_slice1() {
	fruits := make([]string, 3)
	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func Example_slice2() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	// By the way, you cannot use '-' index in Go unlike Python
	// But you can use like fruits[:len(fruits) - 1]
	// Be careful about index of slice. Be. Careful.
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)		// 이어붙이기
	f4 := append(f1[:2], f2...)	// 토마토 빼고 붙이기
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]
}