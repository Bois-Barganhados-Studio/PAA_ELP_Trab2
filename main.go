package main

import (
	"bufio"
	"fmt"
	"main/dividenconquer"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer fmt.Println("FIM")

	// Lê o array de um arquivo
	nums, err := ReadArrayFromFile("inputs/inputs10000.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	arrNum := []int{}
	initTime := time.Now()
	for i := 0; i < 1; i++ {
		arrNum = dividenconquer.MaximumSubarray(nums)
	}
	finalTime := time.Now()
	fmt.Println("Array:", arrNum)
	fmt.Println("Quantidade de comparações:", dividenconquer.Qtd_comp)
	fmt.Println("Tempo de execução:", finalTime.Sub(initTime))
}

func ReadArrayFromFile(filePath string) ([]int, error) {
	// Abre o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Inicializa um scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	// Lê a primeira linha do arquivo
	scanner.Scan()
	fileContent := scanner.Text()

	// Divide a string em um slice de strings usando espaços como delimitador
	numStrings := strings.Fields(fileContent)

	// Inicializa um slice de inteiros para armazenar os valores
	var numbers []int

	// Converte as strings para inteiros e adiciona ao slice
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
