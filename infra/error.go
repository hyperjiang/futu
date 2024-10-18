package infra

import "errors"

// Ret is the interface for response.
type Ret interface {
	GetRetType() int32
	GetRetMsg() string
}

// Error checks if the response is an error.
func Error(r Ret) error {
	if r.GetRetType() != 0 {
		return errors.New(r.GetRetMsg())
	}

	return nil
}
