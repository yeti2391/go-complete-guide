package main

import (
	"fmt"
	"math"
)

func main(){
	const inflacion = 2.5
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

	// Ahora calculo cuanto va a ser el valor a futuro con la formula VF = InvAmount * ((1+ExpectedReturn)/100)^years 
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	fmt.Println(futureValue)
	fmt.Printf("%.2f\n", futureValue)

	// Ahora calculamos el valor real teniendo en cuenta una inflacion constante
	realFutureValue := futureValue / math.Pow(1+inflacion/100,years)
	fmt.Println(realFutureValue)
}