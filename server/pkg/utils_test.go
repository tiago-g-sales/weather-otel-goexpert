package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiago-g-sales/temp-cep/pkg"
)


func TestRemoveAccents(t *testing.T){
	localidade := "Jundiaí"

	expected := "Jundiai"

	result , err := pkg.RemoveAccents(localidade)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	result , err = pkg.RemoveAccents("")
	assert.Nil(t, err)
	expected = ""
	assert.Equal(t, expected, result)
	
}

func TestConvertTemp(t *testing.T){

	temp := 100.0
	expected :=  373.0

	result, err := pkg.ConvertTemp(temp)
	assert.Empty(t, "", result)
	assert.Nil(t, err)
	assert.Equal(t,expected, result )


}

func TestStrigReplace( t *testing.T){
	
	text := "São Paulo"
	expected := "Sao%20Paulo"

	result := pkg.ReplaceAndRemoveAccents(text)
	assert.Equal(t, expected, result)
	assert.NotEqual(t,"São Paulo", result )



}

func BenchmarkRemoveAcentos(b *testing.B) {
	for i := 0 ; i < b.N ; i++{
		pkg.RemoveAccents("São Paulo")
		
	}
	
}