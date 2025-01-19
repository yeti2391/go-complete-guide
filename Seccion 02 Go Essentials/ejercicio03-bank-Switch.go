package main

import "fmt"

func main(){
	var choice int
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your balance is", accountBalance)
		
	case 2:
		fmt.Print("How much do you want to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		//se agrega el nuevo saldo a accountBalance siempre y cuando sea un valor válido
		if depositAmount >= 0{
			accountBalance += depositAmount
			fmt.Println("Balance actualizado! Tu cuenta actual tiene: ", accountBalance)
			
		} else {
			fmt.Println("Valor del deposito no valido")
			
		}
	case 3:
		fmt.Print("How much do you want to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		//si hay saldo suficiente, se realiza la extración y se muestra el nuevo saldo de accountBalance
		if withdrawAmount <= accountBalance && withdrawAmount >= 0{
			accountBalance -= withdrawAmount
			fmt.Println("Balance actualizado! Tu cuenta actualmente tiene: ", accountBalance)
			
		} else if withdrawAmount < 0{ 
			fmt.Println("Valor de extracción no valido")
			
		} else {
				fmt.Println("No tiene suficiente saldo para la extracción!")
				
		}	
	case 4:
		fmt.Printf("Hasta la proxima \n")
		return
	default:
		fmt.Println("Opcion no válida, intenta nuevamente")
		
	}
}

/*
La idea es una aplicación básica que pregunte al usuario que opción realizar:
revisar saldo (check balance),
depositar dinero (deposit monety)
retirar dinero (withdraw money)
exit

*/
