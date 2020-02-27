package json_test

import "time"

import "encoding/json"

import "log"

import "fmt"

type status int
const (
	UNKNOWN	status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

func NewDeadLine(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title		string
	Status		status
	Deadline	*Deadline
}

func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadLine(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	// json package only Marshal the field which starts with UpperCase
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Laundry
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}