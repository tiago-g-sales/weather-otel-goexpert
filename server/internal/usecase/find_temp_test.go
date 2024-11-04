package usecase_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiago-g-sales/temp-cep/pkg"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/mocks"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/usecase"
)



func TestFindTemp(t *testing.T ){

	localidade := "Itapecerica da Serra"

	temperatura := model.Temperatura{
		Temp_C: 24,
		Temp_F: 75.2,
	}

	c := mocks.NewMockClientTemp()
	c.On("FindTemp", localidade,).Return(&model.Temperatura{
		Temp_C: 24.0,
		Temp_F: 75.2,
		Temp_K: 297.0,
	}, nil)

	temp , err := usecase.FindTempHTTPClient.FindTemp(c, localidade)
	if err != nil {
		panic(err)
	}
	resp , err := pkg.ConvertTemp(temp.Temp_C)
	assert.Nil(t, err)
	assert.Equal(t, temp.Temp_C, temperatura.Temp_C )
	assert.Equal(t, temp.Temp_F, temperatura.Temp_F)
	assert.Equal(t, temp.Temp_K, resp)

}



func TestNewHTTPClieTemp(t *testing.T) {
	
	client := usecase.NewHTTPClientTemp(http.Client{})
	assert.NotNil(t, client)

}

