package main

import (
	"github.com/gookit/ini/v2/dotenv"
	"github.com/shivansh98/kahawatein/bootstrap"
	"github.com/shivansh98/kahawatein/utilities"
	"github.com/spf13/viper"
)

func main() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		utilities.CallPanic(err)
	}
	viper.AutomaticEnv()
	bootstrap.InitServices()
}
