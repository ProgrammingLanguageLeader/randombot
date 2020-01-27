package user

import "fmt"

type DoesNotExist struct {
	message string
}

func (err *DoesNotExist) Error() string {
	return fmt.Sprintf("user does not exist: %s", err.message)
}
