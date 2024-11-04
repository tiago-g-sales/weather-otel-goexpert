package usecase_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/mocks"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/usecase"
)


func TestFindCep(t *testing.T ){

	endereco :=  model.ViaCEP{
		Cep: "06865-010",
		Logradouro: "Rua Tancredo de Almeida Neves",
		Complemento: "",
		Bairro: "Jardim do Édem",
		Localidade: "Itapecerica da Serra",
		Uf: "SP",
		Ibge: "3522208",
		Gia: "3700",
		Ddd: "11",
		Siafi: "6545",
	}



	c := mocks.NewMockClientCep()
	c.On("FindCep", endereco.Cep,).Return(&model.ViaCEP {
		Cep: endereco.Cep,
		Logradouro: "Rua Tancredo de Almeida Neves",
		Complemento: "",
		Bairro: "Jardim do Édem",
		Localidade: "Itapecerica da Serra",
		Uf: "SP",
		Ibge: "3522208",
		Gia: "3700",
		Ddd: "11",
		Siafi: "6545",
	}, nil)

	temp , err := usecase.NewHTTPClient(http.Client{}).FindCep(endereco.Cep)	
	assert.NotNil(t, temp, err)
	assert.Equal(t, temp.Cep, endereco.Cep )
	assert.Equal(t, temp.Logradouro, endereco.Logradouro )
	assert.Equal(t, temp.Bairro, endereco.Bairro)
	assert.Equal(t, temp.Complemento, endereco.Complemento)
	assert.Equal(t, temp.Ddd, endereco.Ddd)
	assert.Equal(t, temp.Gia, endereco.Gia)
	assert.Equal(t, temp.Ibge, endereco.Ibge)
	assert.Equal(t, temp.Localidade, endereco.Localidade)
	assert.Equal(t, temp.Uf, endereco.Uf)
	assert.Equal(t, temp.Siafi, endereco.Siafi)

	

}

func TestFindCepEmpty(t *testing.T ){

	
	temp , _ := usecase.NewHTTPClient(http.Client{}).FindCep("")	
	assert.Empty(t, temp)


}


func TestNewHTTPClient(t *testing.T) {
	
	client := usecase.NewHTTPClient(http.Client{})
	assert.NotNil(t, client)

}