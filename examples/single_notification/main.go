/*
 * main.go for go-notify
 * by lenorm_f
 */

package main

import notify "github.com/mqu/go-notify"
import (
	"os"
	"fmt"
	"time"
)

const (
	DELAY = 3000;
)

func main() {
	notify.Init("Hello World!")
	hello := notify.NotificationNew("Hello World!",
		"This is an example notification.",
		"")

	if hello == nil {
		fmt.Fprintf(os.Stderr, "Unable to create a new notification\n")
		return
	}
	// hello.SetTimeout(3000)
	notify.NotificationSetTimeout(hello, DELAY)

	// hello.Show()
	if e := notify.NotificationShow(hello); e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e.Message())
		return
	}

	time.Sleep(DELAY * 1000000)
	// hello.Close()
	notify.NotificationClose(hello)

	notify.UnInit()
}
