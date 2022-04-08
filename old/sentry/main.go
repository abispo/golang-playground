package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"

	"github.com/getsentry/sentry-go"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.Get("DSN").(string),
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
