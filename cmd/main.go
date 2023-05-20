package main

import (
	"fmt"
	"sth"
	"sth/handler"
	"sth/repository"
	"sth/service"

	"github.com/spf13/viper"
)

func main() {
	initConfig()
	db := repository.NewPostgresDB(repository.Config{
		Addr:     viper.GetString("DBADDR"),
		User:     viper.GetString("DBUSER"),
		Password: viper.GetString("DBPASSWORD"),
		Database: viper.GetString("DB"),
	})
	fmt.Println(db)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(sth.Server)
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			panic(err)
		}

}


func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	res := viper.ReadInConfig()
	if res != nil {
		panic(res)
	}
	return 
}