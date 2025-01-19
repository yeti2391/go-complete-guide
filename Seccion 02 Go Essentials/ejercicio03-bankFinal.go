package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)
const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	//se lee el archivo balance.txt y se almacena una colección de bytes ([]bytes) en la variable balanceText
	data, err := os.ReadFile(accountBalanceFile)
	// si queremos manejar el error:
	if err != nil {
		return 1000, errors.New("No se pudo leer el archivo balance.txt")
	} // lo que decimos que si hay un error (por ejemplo no existe el archivo balance.txt), el saldo será de 1000, y de no haber entonces err=nil y se sigue con el código

	//como os.ReadFile lo que lee es una colección de bytes lo que se hace es convertirlo a string
	balanceText := string(data)
	//y despues se convierte el string a float64
	//la función strconv.ParseFloat recibe dos parametros, el primero es el string que se quiere convertir y el segundo es el bitSize
	//y devuelve dos valores, el primero es el valor convertido y el segundo es un error que si no estamos interesados en su manejo lo ignoramos con _
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("No se pudo convertir el saldo a un número")
	}

	return balance, nil 
	//este ultimo return sería el caso donde estuviera todo ok, como ahora la función devuelve dos valores 
	//se indica que devuelva el balance y nil, que es un valor nulo que indica que no hay error


}

func writeBalanceToFile(balance float64) {
	// Lo primero que se hace es convertir el parametro balance a una string
	balanceText := fmt.Sprint(balance)
	//a continuación usamos la función WriteFile de la librería os para escribir el archivo balance.txt
	//el primer parametro es el nombre del archivo, el segundo es el contenido del archivo y el tercero es el permiso
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// 0644 es un permiso de lectura y escritura, lo que significa que el archivo puede ser leído y escrito por el propietario,
	// pero solo leído por otros usuarios.

}

func main(){
	var accountBalance = getBalanceFromFile() 
	var choice int


	fmt.Println("Welcome to Go Bank!")
	for{
		
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much do you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			//se agrega el nuevo saldo a accountBalance siempre y cuando sea un valor válido
			if depositAmount >= 0{
				accountBalance += depositAmount
				fmt.Println("Balance actualizado! Tu cuenta actual tiene: ", accountBalance)
				writeBalanceToFile(depositAmount)
				continue
			} else {
				fmt.Println("Valor del deposito no valido")
				continue
			}
			
		} else if choice == 3 {
			fmt.Print("How much do you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			//si hay saldo suficiente, se realiza la extración y se muestra el nuevo saldo de accountBalance
			if withdrawAmount <= accountBalance && withdrawAmount >= 0{
				accountBalance -= withdrawAmount
				fmt.Println("Balance actualizado! Tu cuenta actualmente tiene: ", accountBalance)
				writeBalanceToFile(withdrawAmount)
				continue
			} else if withdrawAmount < 0{ 
				fmt.Println("Valor de extracción no valido")
				continue
			} else {
					fmt.Println("No tiene suficiente saldo para la extracción!")
					continue
			}	

		} else if choice == 4{
			fmt.Printf("Hasta la proxima \n")
			break
		} else {
			fmt.Println("Opcion no válida, intenta nuevamente")
			continue
		}
	}


}

/*
La idea es una aplicación básica que pregunte al usuario que opción realizar:
revisar saldo (check balance),
depositar dinero (deposit monety)
retirar dinero (withdraw money)
exit

*/
