package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiago-g-sales/weather-otel-goexpert/internal/usecase"
)

const(
	INVALID_ZIP_CODE = "invalid zipcode"
	CAN_NOT_FIND_ZIPCODE = "can not find zipcode"
	QUERY_PARAMETER = "cep"
	LEN_ZIP_CODE = 8
)


func FindTempByCepHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get(QUERY_PARAMETER)
	if cepParam == ""{
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(cepParam) > LEN_ZIP_CODE || len(cepParam) < LEN_ZIP_CODE  {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(INVALID_ZIP_CODE)
		return		
	}

	clientCep := usecase.NewHTTPClient(http.Client{})

	cep, err :=usecase.FindCepHTTPClient.FindCep(clientCep, cepParam)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(CAN_NOT_FIND_ZIPCODE)
		return
	}

 	clientTemp := usecase.NewHTTPClientTemp(http.Client{})

	temp , err := usecase.FindTempHTTPClient.FindTemp(clientTemp, cep.Localidade)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)

}