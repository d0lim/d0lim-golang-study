package sort_test

import "strings"

import "sort"

import "fmt"

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string {
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}