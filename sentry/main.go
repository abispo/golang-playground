package main

import (
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "http://c0d4f983c6094b14961f720cf36a60a9@127.0.0.1:9000/2",
	})

	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	f, err := os.Open("filename.ext")
	fmt.Println(f)
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(time.Second * 5)
	}
}
