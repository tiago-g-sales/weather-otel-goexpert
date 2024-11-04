package pkg



func ConvertTemp( tempC float64 ) (float64, error){

	const K = 273
	return tempC + K, nil

}



