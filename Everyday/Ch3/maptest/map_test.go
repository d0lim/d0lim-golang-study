package maptest

import (
	"reflect"
	"testing"
)

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}