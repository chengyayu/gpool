package wraper

import (
	"fmt"
	"runtime"
)

func RecoveredFn(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				err = fmt.Errorf("recoverfn: panic recovered: %s\n%s", r, buf)
			}
		}()

		return f()
	}
}
