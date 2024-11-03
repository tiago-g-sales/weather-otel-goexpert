package handler_test

import (

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiago-g-sales/temp-cep/internal/handler"
	"github.com/tiago-g-sales/temp-cep/internal/mocks"
	"github.com/tiago-g-sales/temp-cep/internal/model"
	"github.com/tiago-g-sales/temp-cep/internal/usecase"
)

const (
	INVALID_ZIP_CODE = "invalid zipcode"
	CAN_NOT_FIND_ZIPCODE = "can not find zipcode"
	QUERY_PARAMETER = "cep"
	LEN_ZIP_CODE = 8
)


func TestFindTempByCep(t *testing.T) {
	cep := "Jundiai"
	localidade := "Jundiai"

	temperatura := &model.Temperatura{
		Temp_C: 0,
		Temp_F: 0,
		Temp_K: 0,
	}	

	ct := mocks.NewMockClientTemp()
	ct.On("FindTemp", localidade).Return(&model.Temperatura{}, nil)

	temp, err := usecase.FindTempHTTPClient.FindTemp( ct, cep)
	assert.Nil(t, err)
	assert.Equal(t, temp.Temp_C, temperatura.Temp_C)
	assert.Equal(t, temp.Temp_F, temperatura.Temp_F)
	assert.Equal(t, temp.Temp_K, temperatura.Temp_K)

	}

func TestFindTempByCepHandler_InvalidZipCode(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.FindTempByCepHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Contains(t, rr.Body.String(), INVALID_ZIP_CODE)
}







