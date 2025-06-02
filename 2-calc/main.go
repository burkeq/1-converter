package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func startCalc(){
	operationsVariants := [3]string{"AVG", "SUM", "MED"}
	fmt.Println("___ Calculator ___")
	var operation string
	fmt.Printf("Введите операцию(%v): ", strings.Join(operationsVariants[:], " "))
	fmt.Scan(&operation)
	var userNumbers string
	fmt.Print("Введите числа через запятую(1,2,3): ")
	fmt.Scan(&userNumbers)
	numbers, err := handlingUserNumbers(userNumbers)
	if err != nil{
		fmt.Println("Ошибка: Неверный формат ввода, попробуйте без пробела!")
	}
	calcResult, err := calcOperation(strings.ToUpper(operation), numbers)
	if err != nil{
		fmt.Printf("Ошибка: %v\n", err.Error())
	}
	fmt.Printf("Ваш резуьтат: %v", calcResult)
}

func calcOperation(operation string, num []int) (int, error){
	switch operation{
	case "AVG":
		return calcAVG(num), nil
	case "SUM":
		return calcSum(num), nil
	case "MED":
		return calcMED(num), nil
	default:
		return 0, errors.New("Неверный тип операции")
	}
}

func calcSum(nums []int) int{
	var result int
	for _, value := range nums{
		result += value
	} 
	return result
}

func calcAVG(nums []int) int{
	sumOfNums := calcSum(nums)
	return sumOfNums / len(nums)
}

func calcMED(nums []int) int{
	return nums[(len(nums)-1)/2]
}

func handlingUserNumbers(stringNum string)([]int, error){
	arrStrNum := strings.Split(stringNum, ",")
	numbers := make([]int, 0, len(arrStrNum))
	for _, s := range arrStrNum{
		num, err := strconv.Atoi(s)
		if err != nil{
			return nil, errors.New(err.Error())
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}