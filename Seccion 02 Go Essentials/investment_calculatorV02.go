package main

import (
	"fmt"
	"math"
)

//se reescribe el código utilizando funciones
const inflacion = 2.5

func main(){
	
	var investmentAmount,	years,	expectedReturnRate float64

	// solicito que el usuario ingrese el monto a invertir
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// solicito por cuantos años sera la inversion
	fmt.Print("Years: ")
	fmt.Scan(&years)

	// solicito las ganancias esperadas
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//obtenidos los valores realizamos los calculos en la func calculando y los guardamos en 2 variables
	futureValue, realFutureValue := calculando(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Valor futuro: %.2f\nEl valor real es: %.2f", futureValue, realFutureValue)
}


func calculando(inversion, tasaRetorno, tiempo float64) (float64, float64){
	// Ahora calculo cuanto va a ser el valor a futuro con la formula VF = InvAmount * ((1+ExpectedReturn)/100)^years 
	fv := inversion * math.Pow(1+tasaRetorno/100,tiempo)
	
	// Ahora calculamos el valor real teniendo en cuenta una inflacion constante
	rfv := fv / math.Pow(1+inflacion/100,tiempo)
	
	//damos como salida ambas operaciones
	return fv, rfv
}