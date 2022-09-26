package bulk

import (
	"fmt"
	"strings"
)

type ErrBulkEstimate map[int]*ErrEstimate

func (e ErrBulkEstimate) Error() string {
	msgs := []string{}
	for k, v := range e {
		msg := fmt.Sprintf("%v: %v", k, v.Error())
		msgs = append(msgs, msg)
	}
	return strings.Join(msgs, "\n")
}

type ErrEstimate struct {
	Inner error
}

func (e ErrEstimate) Error() string {
	return fmt.Sprintf("estimate error: %v", e.Inner.Error())
}
