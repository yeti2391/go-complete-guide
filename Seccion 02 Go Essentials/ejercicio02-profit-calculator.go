package main

import "fmt"

func main(){
	var revenue, expenses, tax_rate float64

	// datos a solicitar al usuario
	fmt.Print("Ingrese los ingresos: ")
	fmt.Scan(&revenue)
	fmt.Print("Ingrese los costos: ")
	fmt.Scan(&expenses)
	fmt.Print("Ingrese los impuestos: ")
	fmt.Scan(&tax_rate)

	//calculo de ganancias sin impuestos
	ebt := revenue-expenses
	fmt.Printf("Las ganancias antes de impuestos (EBT) es de: %.2f\n", ebt)

	//calculo de ganancias después de los ingresos
	impuestos := ebt * (tax_rate/100)
	profit := ebt - impuestos
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