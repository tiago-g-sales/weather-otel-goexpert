package usecase

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tiago-g-sales/temp-cep/configs"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"

	"github.com/tiago-g-sales/temp-cep/pkg"
	"github.com/valyala/fastjson"
)

type FindTempHTTPClient interface {
    FindTemp(loc string) (*model.Temperatura, error)
}

type HTTPClientTemp struct {
    client *http.Client
}

func NewHTTPClientTemp (client http.Client) (*HTTPClientTemp){
	return &HTTPClientTemp{client: &client}
}


func (h *HTTPClientTemp) FindTemp(localidade string) (*model.Temperatura, error) {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	temp := model.Temperatura{}
	temp.City = localidade

	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query() 
	q.Add("q", localidade)
	q.Add("key", configs.API_KEY )
	req.URL.RawQuery = q.Encode()		

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p fastjson.Parser

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	v, err := p.ParseBytes(body)
	if err != nil {
		panic(err)
	}


	json.Unmarshal([]byte(v.GetObject("current").String()), &temp)

	temp.Temp_K, _ = pkg.ConvertTemp(temp.Temp_C)

	return &temp, nil
}
