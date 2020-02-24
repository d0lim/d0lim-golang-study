package fix_arg_test

import (
	"bufio"
	"fmt"
	"io"
)

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

// Insert 함수는 집합에 val을 추가한다.
func Insert(m MultiSet, val string) {
	fmt.Println(val, " inserted")
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func NewMultiSet() MultiSet{
	return MultiSet{}
}

func ExampleNotExecutable() {
	m := NewMultiSet()
	ReadFrom(r, func(line string) {
		Insert(m, line)
	})
}

func InsertFunc(m MultiSet) func(val string) {
	// 클로저를 만들고, m에만 Insert하는 함수를 반환한다.
	return func(val string) {
		Insert(m, val)
	}
}

func ExampleNotExecutable_abstract_first() {
	m := NewMultiSet()
	ReadFrom(r, InsertFunc(m))
}

func BindMap(f SetOp, m MultiSet) func(val string) {
	// InsertFunc의 더 일반화된 버전이다. SetOperator 타입의 함수 f를 받아서, m에만 f하는 함수를 반환한다.
	return func(val string) {
		f(m, val)
	}
}

func ExampleNotExecutable_abstract_second() {
	m := NewMultiSet()
	ReadFrom(r, BindMap(Insert, m))
}