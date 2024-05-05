package errors

import (
	"errors"
	"fmt"
)

func Wrap(err error, msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", msg, err))
}
