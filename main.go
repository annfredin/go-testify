package main

import (
	"fmt"
	"go-testing/common"
)

func main(){

	fmt.Println("Testing Demo!!!")

	fac := common.CalFactorial(5)
	fmt.Println("Factorial => ", fac)
}