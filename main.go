package main

import (
	"context"
	"github.com/gookit/ini/v2/dotenv"
	"github.com/shivansh98/kahawatein/internal/bootstrap"
	"github.com/shivansh98/kahawatein/internal/services"
	"github.com/shivansh98/kahawatein/utilities"
	"github.com/spf13/viper"
	"os/signal"
	"syscall"
)

func main() {
	err := dotenv.Load("./", ".env")
	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGTERM, syscall.SIGKILL)
	if err != nil {
		utilities.CallPanic(err)
	}
	viper.AutomaticEnv()
	bootstrap.InitServices()
	srv := services.InitHTTPServer()

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			utilities.Logger.Println("Errror")
		}
	}()

	<-ctx.Done()
	utilities.Logger.Println("Signal received")
	err = srv.Shutdown(ctx)
	if err != nil {
		utilities.Logger.Println(err)
	}
	cancel()
}
