package main

import (
	"github.com/spf13/viper"
	kraicklistHttp "github.com/zulkan/kraicklist/kraicklist/delivery/http"
	"github.com/zulkan/kraicklist/kraicklist/repository/file"
	kraicklistUseCase "github.com/zulkan/kraicklist/kraicklist/usecase"
	"strconv"

	"fmt"
	"log"
	"os"
)

func main() {
	//init viper to read env
	viper := viper.New()
	viper.BindEnv("FILE_DATA")
	viper.SetDefault("FILE_DATA", "data.gz")
	pathData := viper.GetString("FILE_DATA")
	fmt.Println("PATH_DATA", pathData)

	// initialize searcher
	repo, err := file.NewFileRecordRepository(pathData)
	if err != nil {
		log.Fatalf("unable to load search data due: %v", err)
	}
	// define http handlers

	// define port, we need to set it as env for Heroku deployment
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3001
	}

	service := kraicklistUseCase.NewSearcherUseCase(repo)
	// start server
	fmt.Printf("Server is listening on %s...", port)
	err = kraicklistHttp.NewHttpServer(port, service)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}

