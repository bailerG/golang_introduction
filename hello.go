package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	urls := openFile()

	for i := range urls {
		request(urls[i])
	}
	printLogs()
	clearLogs()
}

func request(site string) {
	result, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro com a requisição: ", err)
		return
	}

	fmt.Println("Resultado: ", result.Status)

	if result.StatusCode == 200 {
		createLog(site, true)
		return
	} else {
		createLog(site, false)
		return
	}
}

func openFile() []string {
	var result []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		result = append(result, line)

		if err == io.EOF {
			break
		}
	}

	fileErr := file.Close()

	if fileErr != nil {
		fmt.Println("Erro ao fechar o arquivo: ", fileErr)
	}

	return result
}

func createLog(text string, status bool) {
	file, err := os.OpenFile("loggin.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erro ao fechar arquivo: ", err)
		}
	}(file)

	if err != nil {
		fmt.Println("Erro ao registrar logs: ", err)
	}

	_, err = file.WriteString(time.Now().Format("02/01/2006 - 15:04:05 ") + text + " - online: " + strconv.FormatBool(status) + "\n")
	if err != nil {
		return
	}

}

func printLogs() {
	file, err := os.ReadFile("loggin.txt")
	if err != nil {
		fmt.Println("Erro ao abrir logs: ", err)
		return
	}

	fmt.Println(string(file))
}

func clearLogs() {
	err := os.Remove("loggin.txt")
	if err != nil {
		fmt.Print("Erro ao excluir logs: ", err)
	}
	return
}
