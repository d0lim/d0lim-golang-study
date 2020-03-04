package main // import "github.com/imdigo/DolimGoLangStudy/Everyday/Ch6/taskmanfirst"

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

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

// MemoryDataAccess is a simple in-memory database.
type MemoryDataAccess struct {
	tasks	map[ID]task.Task
	nextID	int64
}

// NewMemoryDataAccess returns a new MemoryDataAccess.
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess {
		tasks: map[ID]task.Task{},
		nextID: int64(1),
	}
}

//ErrTaskNotExist occurs when the task with the ID was not found
var ErrTaskNotExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put updates a task with a given ID with it.
func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new task.
func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

// ResponeError is the error for the JSON Response.
type ResponeError struct {
	Err error
}

// MarshalJSON returns the JSON representation of the error.
func (err ResponeError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

// UnmarshalJSON parses the JSON representation of the error.
func (err *ResponeError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNotExist.Error() {
			err.Err = ErrTaskNotExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

// Response is a struct for the JSON response.
type Response struct {
	ID		ID				`json:"id,omitempty"`
	Task	task.Task		`json:"task"`
	Error	ResponeError	`json:"error"`
}

// FIXME: m is NOT thread-safe
var m = NewMemoryDataAccess()

const pathPrefix = "/api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}
	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t task.Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}

	switch r.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response {
			ID:		id,
			Task: 	t,
			Error:	ResponeError{err},
		})
		if err != nil {
			log.Println(err)
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(w).Encode(Response {
				ID:		id,
				Task: 	t,
				Error:	ResponeError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "POST":
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(w).Encode(Response {
				ID:		id,
				Task: 	t,
				Error:	ResponeError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response {
			ID:		id,
			Error:	ResponeError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

}


func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}

