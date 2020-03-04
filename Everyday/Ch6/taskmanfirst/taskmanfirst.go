package taskmanfirst // import "github.com/imdigo/DolimGoLangStudy/Everyday/Ch6/taskmanfirst"

import (
	"fmt"

	"github.com/imdigo/DolimGoLangStudy/task"
)

// ID is a data type to identify a task.
type ID string

// DataAccess is an interface to access tasks.
type DataAccess interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error 
}

func main() {
	fmt.Println("Hello World!")
}

