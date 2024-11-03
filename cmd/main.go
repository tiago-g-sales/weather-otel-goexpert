package main

import (
	"net/http"

	"github.com/tiago-g-sales/weather-otel-goexpert/configs"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/handler"
)



func main(){

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	http.HandleFunc("/", handler.FindTempByCepHandler)
	http.ListenAndServe(configs.WebServerPort, nil)
}








