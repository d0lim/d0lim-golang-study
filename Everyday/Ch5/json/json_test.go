package json_test

import "time"

// import "encoding/json"

// import "log"

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
	Title		string		`json:"title,omitempty"`
	Status		status		`json:"status,omitempty"`
	Deadline	*Deadline	`json:"deadline,omitempty"`
	Priority	int			`json:"priority,omitempty"`
	SubTasks	[]Task		`json:"subTasks,omitempty"`
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix + " ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

// func Example_marshalJSON() {
// 	t := Task{
// 		"Laundry",
// 		DONE,
// 		NewDeadLine(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
// 	}
// 	// json package only Marshal the field which starts with UpperCase
// 	b, err := json.Marshal(t)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(string(b))
// 	// Output:
// 	// {"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}
// }

// func Example_unmarshalJSON() {
// 	b := []byte(`{"Title":"Laundry","Status":2,"Deadline":"2015-08-16T15:43:00Z"}`)
// 	t := Task{}
// 	err := json.Unmarshal(b, &t)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(t.Title)
// 	fmt.Println(t.Status)
// 	fmt.Println(t.Deadline.UTC())
// 	// Output:
// 	// Laundry
// 	// 2
// 	// 2015-08-16 15:43:00 +0000 UTC
// }

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title: "Laundry",
		Status: TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title: "Wash",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title: "Dry",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title: "Fold",
			Status: TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	// Output:
	// [ ] Laundry <nil>
	//  [ ] Wash <nil>
	//   [v] Put <nil>
	//   [ ] Detergent <nil>
	//  [ ] Dry <nil>
	//  [ ] Fold <nil>
}