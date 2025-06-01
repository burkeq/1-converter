package main

import (
	"fmt"
	"strings"
)

func main(){
	var (
		selectedSourceCurrency string
		selectedTargetCurrency string
		amountCurrency int
	)
	fmt.Println("___ Конвертер валюты ___")
	selectedSourceCurrency, _ = openUserInput("source")
	_, amountCurrency = openUserInput("amount")
	selectedTargetCurrency, _ = openUserInput("target")

	fmt.Printf("Теперь у вас: %.0f %v", convert(amountCurrency, selectedSourceCurrency, selectedTargetCurrency), strings.ToUpper(selectedTargetCurrency))
}

func convert(amount int, sourceCurrency, targetCurrency string) float64{
	const USDtoEUR = 0.88
	const USDtoRUB = 78.63
	const EURtoRUB = USDtoRUB / USDtoEUR
	rates := map[string]float64 {
		"USD": 1,
		"EUR": USDtoEUR,
		"RUB": USDtoRUB,
	}
	amountInUSD := float64(amount) / rates[strings.ToUpper(sourceCurrency)]
	return amountInUSD * rates[strings.ToUpper(targetCurrency)]
}

func checkCurrency(currency string, currencyVariants [3]string) bool{
	upperSelected := strings.ToUpper(currency)
	if upperSelected == currencyVariants[0] || upperSelected == currencyVariants[1] || upperSelected == currencyVariants[2]{
		return true
	}
	return false
}


func openUserInput(inputType string) (string, int){
	var (
		selectedCurrency string
		amountCurrency int
	)
	currencyVariants := [3]string{"USD", "EUR", "RUB"}
	if inputType == "amount"{
		for{
			fmt.Print("Выбирите количество валюты: ")
			fmt.Scan(&amountCurrency)

			if amountCurrency > 0{
				break
			}
			continue
		}
		return "", amountCurrency
	}
	var currencyType string
	switch inputType{
		case "source":
			currencyType = "исходную"
		case "target":
			currencyType = "целевую"
	}
	for{
		fmt.Printf("Выбирите %v валюту(%v): ", currencyType, strings.Join(currencyVariants[:], " "))
		fmt.Scan(&selectedCurrency)
		if checkCurrency(selectedCurrency, currencyVariants){break}
		continue
	}
	return selectedCurrency, 0
}	