package errors

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func Middleware(handles ...func(*message.Message, error) error) message.HandlerMiddleware {
	var handle func(*message.Message, error) error
	if len(handles) == 0 {
		handle = func(msg *message.Message, err error) error { return err }
	} else {
		handle = handles[0]
	}
	return func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			msgs, err := h(msg)
			err = handle(msg, err)
			return msgs, err
		}
	}
}
