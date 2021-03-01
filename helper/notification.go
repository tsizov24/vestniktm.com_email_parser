package helper

import (
	"fmt"
)

func notFatalNotification(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
