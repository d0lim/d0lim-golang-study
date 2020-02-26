package main

import "fmt"

import "time"

// CountDown is blocking timer function of 'seconds' seconds
func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("2 Seconds left! I'm so excited!!")
	})
	fmt.Println("Ladies and gentlemen!")
	CountDown(5)
}