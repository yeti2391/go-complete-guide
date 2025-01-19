package main

import {
	"fmt"
	"os"
}

func writeBalanceToFile(balance float64) {
	// Lo primero que se hace es convertir el parametro balance a una string
	balanceText := fmt.Sprint(balance)
	//a continuación usamos la función WriteFile de la librería os para escribir el archivo balance.txt
	//el primer parametro es el nombre del archivo, el segundo es el contenido del archivo y el tercero es el permiso
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
	// 0644 es un permiso de lectura y escritura, lo que significa que el archivo puede ser leído y escrito por el propietario,
	// pero solo leído por otros usuarios.

}

func main(){
	var choice int
	var accountBalance float64 = 1000

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
