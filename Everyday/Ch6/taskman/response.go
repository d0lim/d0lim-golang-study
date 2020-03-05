package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/imdigo/DolimGoLangStudy/task"
)

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
		if tv == task.ErrTaskNotExist.Error() {
			err.Err = task.ErrTaskNotExist
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
	ID		task.ID				`json:"id,omitempty"`
	Task	task.Task		`json:"task"`
	Error	ResponeError	`json:"error"`
}