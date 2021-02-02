package skylib

import (
	"context"
)

//CheckDeadline function
func CheckDeadline(ctx context.Context, funcName string, ret ...interface{}) interface{} {
	if ctx.Err() == context.DeadlineExceeded {
		Infof("%v function is DeadlineExceeded", funcName)
		return ret
	}
	return nil
}
