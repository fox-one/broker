package broker

import (
	"fmt"
)

type Err struct {
	code int
	msg  string
}

func (e Err) Error() string {
	return fmt.Sprintf("%d %s", e.code, e.msg)
}
