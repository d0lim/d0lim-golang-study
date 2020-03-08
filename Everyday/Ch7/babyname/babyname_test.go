package babyname

import "fmt"

func ExampleBabyNames() {
	for n := range BabyNames("성정명재경", "준호우훈진") {
		fmt.Println(n)
	}
	// Output:
	// 성준
	// 성호
	// 성우
	// 성훈
	// 성진
	// 정준
	// 정호
	// 정우
	// 정훈
	// 정진
	// 명준
	// 명호
	// 명우
	// 명훈
	// 명진
	// 재준
	// 재호
	// 재우
	// 재훈
	// 재진
	// 경준
	// 경호
	// 경우
	// 경훈
	// 경진
}