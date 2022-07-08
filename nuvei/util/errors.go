package util

import "fmt"

type NuveiError struct {
	Err error
}

func (r *NuveiError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}
