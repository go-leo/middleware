package recovery

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
)

func Middleware(handles ...func(*message.Message, any) error) message.HandlerMiddleware {
	var handle func(*message.Message, any) error
	if len(handles) == 0 {
		handle = func(msg *message.Message, p any) (err error) {
			return fmt.Errorf("panic triggered: %+v", err)
		}
	} else {
		handle = handles[0]
	}
	return func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) (msgs []*message.Message, err error) {
			panicked := true

			defer func() {
				if p := recover(); p != nil || panicked {
					err = handle(msg, p)
				}
			}()

			msgs, err = h(msg)
			panicked = false
			return msgs, err
		}
	}
}
