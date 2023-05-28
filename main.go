package main

import (
	"context"
	"github.com/gookit/ini/v2/dotenv"
	"github.com/shivansh98/kahawatein/internal/bootstrap"
	"github.com/shivansh98/kahawatein/internal/services"
	"github.com/shivansh98/kahawatein/utilities"
	"github.com/spf13/viper"
	"sync"
	"time"
)

func main() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		utilities.CallPanic(err)
	}
	viper.AutomaticEnv()
	bootstrap.InitServices()
	ctx, cancel := context.WithCancel(context.TODO())
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		services.InitHTTPServer(ctx)
		wg.Done()
	}()
	time.Sleep(1 * time.Second)
	wg.Wait()
	cancel()
}
