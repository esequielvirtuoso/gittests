package utils

import "fmt"

func PrintVars(varName string, varValue interface{}){
	fmt.Printf("Variável: %s",varName)
	fmt.Println("")
	fmt.Println("Valor: ",varValue)
}
