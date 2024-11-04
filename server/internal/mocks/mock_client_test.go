package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"
)

func TestMockClientCep_FindCep(t *testing.T) {
	mockCepClient := new(mockClientCep)
	expectedCep := &model.ViaCEP{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:     "Sé",
		Localidade: "São Paulo",

	}

	mockCepClient.On("FindCep", "01001-000").Return(expectedCep, nil)

	cep, err := mockCepClient.FindCep("01001-000")

	assert.NoError(t, err)
	assert.Equal(t, expectedCep, cep)
	mockCepClient.AssertExpectations(t)
}

func TestMockClientCep_FindCep_Error(t *testing.T) {
	mockCepClient := new(mockClientCep)
	mockCepClient.On("FindCep", "00000-000").Return(nil, assert.AnError)

	cep, err := mockCepClient.FindCep("00000-000")

	assert.Error(t, err)
	assert.Nil(t, cep)
	mockCepClient.AssertExpectations(t)
}

func TestMockClientTemp_FindTemp(t *testing.T) {
	mockTempClient := new(mockClientTemp)
	expectedTemp := &model.Temperatura{
		Temp_C: 25.0,
	}

	mockTempClient.On("FindTemp", "São Paulo").Return(expectedTemp, nil)

	temp, err := mockTempClient.FindTemp("São Paulo")

	assert.NoError(t, err)
	assert.Equal(t, expectedTemp, temp)
	mockTempClient.AssertExpectations(t)
}

func TestMockClientTemp_FindTemp_Error(t *testing.T) {
	mockTempClient := new(mockClientTemp)
	mockTempClient.On("FindTemp", "Unknown").Return(nil, assert.AnError)

	temp, err := mockTempClient.FindTemp("Unknown")

	assert.Error(t, err)
	assert.Nil(t, temp)
	mockTempClient.AssertExpectations(t)
}