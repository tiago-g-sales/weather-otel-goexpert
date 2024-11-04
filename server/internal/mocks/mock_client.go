package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"
)


type mockClientCep struct { mock.Mock }
type mockClientTemp struct { mock.Mock }

func NewMockClientCep() *mockClientCep {
	 return &mockClientCep{} 
	}


func NewMockClientTemp() *mockClientTemp {
	return &mockClientTemp{} 
	}
   


func (c *mockClientCep) FindCep(cep string) (*model.ViaCEP, error){
    
	args := c.Called(cep)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*model.ViaCEP), args.Error(1)
}

func (c *mockClientTemp) FindTemp(loc string) (*model.Temperatura, error){
    
	args := c.Called(loc)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*model.Temperatura), args.Error(1)
}