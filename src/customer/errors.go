package customer

import (
	"fmt"
)

type NotFoundError struct {
	message  string
	HttpCode int
}

type ElasticsearchConnectionError struct {
	OriginalError error
	HttpCode      int
}

func (err NotFoundError) Error() string {
	return err.message
}

func (err ElasticsearchConnectionError) Error() string {
	return fmt.Sprintf("There is a connection error with elasticsearch - (%+v)", err.OriginalError)
}
