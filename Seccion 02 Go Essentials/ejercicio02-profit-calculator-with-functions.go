package main

import "fmt"

//se re hace utilizando funciones

func main(){
	// datos a solicitar al usuario
	revenue := obtenerDatos("Ingresos: ") 
	expenses := obtenerDatos("Gastos: ")
	tax_rate := obtenerDatos("Impuestos: ")

	calculando(revenue, expenses, tax_rate) 
}
 
//se crea una única función para pedir input del usuario
func obtenerDatos(infoText string) float64{
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
 }


//se crea una única función para realizar los calculos e imprimir resultados
func calculando(userRevenue, userExpenses, userTax_rate float64){
	//calculo de ganancias sin impuestos
	ebt := userRevenue-userExpenses
	fmt.Printf("Las ganancias antes de impuestos (EBT) es de: %.2f\n", ebt)

	//calculo de ganancias después de los ingresos
	profit := ebt * (1 - userTax_rate / 100)
	fmt.Printf("Después de los impuestos las ganancias son: %.2f\n",profit)

	//calculo de ratio
	ratio := ebt/profit
	fmt.Printf("El calculo de ratio es de: %.2f\n", ratio)
 }


/*
"calculadora de ganancias" en Go. 
Esta calculadora tomará como entrada los ingresos (revenue), los gastos (expenses) y la tasa de impuestos (tax rate), 
y calculará las ganancias antes de impuestos (EBT) y las ganancias después de impuestos (profit), y el ratio (EBT / profit)

*/
