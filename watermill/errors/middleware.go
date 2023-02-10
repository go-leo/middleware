package errors

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func Middleware(errorHandler func(err error) error) message.HandlerMiddleware {
	if errorHandler == nil {
		errorHandler = func(err error) error {
			return err
		}
	}
	return func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			msgs, err := h(msg)
			err = errorHandler(err)
			return msgs, err
		}
	}
}
