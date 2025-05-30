package main

import "fmt"

func main(){
	const USDtoEUR = 0.88
	const USDtoRUB = 78.63
	const EURtoRUB = USDtoEUR / USDtoRUB
	fmt.Println("___ Конвертер валют ___")
	fmt.Println(getUserData())
	calcCurrencyToCurrency(1, "asd", "asd")
}

func getUserData() int{
	var amountCurrency int
	fmt.Print("Введите количество валюты: ")
	fmt.Scan(&amountCurrency)
	return amountCurrency
}

func calcCurrencyToCurrency(amountCurrency int, currency1, currency2 string) int{
	return 0
}