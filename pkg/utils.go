package pkg

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)



func RemoveAccents(s string) (string, error)  {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		return "", e
	}
	return output, nil
}


func ConvertTemp( tempC float64 ) (float64, error){

	const K = 273
	return tempC + K, nil

}

func ReplaceAndRemoveAccents(s string) (string){
	
	r, _ := RemoveAccents(s)
	
	return strings.Replace(r, " ", "%20", -1)

}

