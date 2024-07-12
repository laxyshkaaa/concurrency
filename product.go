package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	bytes := []string{"Тинькофф", "Прибыль", "Кредит", "Банк", "Налог"}

	file, err := os.ReadFile("Отчет.txt")

	if err != nil{
		log.Fatal(err.Error())
		return
	} 
	for _, val := range bytes{
		wordcount := strings.Count(string(file), val)
		fmt.Println("слово ", val, " встречается ", wordcount, " раз")
	}
}