package functions

import (
	"fmt"
	"time"
)

func MiddlewareLOG(f func(string)) func(string) {
	return func(name string) {
		fmt.Println(time.Now())
		f(name)
	}
}
