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
	/// Getting the urls and storing on a variable
	urls := openFile()

	/// Creating an iteration through all urls and creating a request with each one
	for i := range urls {
		request(urls[i])
	}

	/// Shows the logs
	printLogs()

	/// Clears the logs
	clearLogs()
}

// / Creates a request on a given url and treats the return
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

// / Opens the urls.txt file where all urls are stored and returns the urls to be used
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

// / Creates a log file to store the request results
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

// / Shows logs on the terminal
func printLogs() {
	file, err := os.ReadFile("loggin.txt")
	if err != nil {
		fmt.Println("Erro ao abrir logs: ", err)
		return
	}

	fmt.Println(string(file))
}

// / Deletes the logs file
func clearLogs() {
	err := os.Remove("loggin.txt")
	if err != nil {
		fmt.Print("Erro ao excluir logs: ", err)
	}
	return
}
