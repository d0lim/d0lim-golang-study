package deadline_test

import "time"

import "fmt"

type Deadline time.Time
type status int
const (
	UNKNOWN	status = iota
	TODO
	DONE
)

type Task struct {
	Title		string
	Status	 	status
	Deadline	*Deadline
}

// OverDue returns true if the deadline is before the current time
func (d *Deadline) OverDue() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}

// OverDue returns true if the deadline is before the current time
func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}

func ExampleDeadline_OverDue() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"no due", TODO, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

