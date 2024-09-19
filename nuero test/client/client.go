package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Response string `json:"response"`
}

func main() {
	url := "http://localhost:5000/ask"

	var userInput string
	fmt.Print("Введите ваш вопрос: ")
	fmt.Scanln(&userInput)

	requestBody := Request{Text: userInput}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Ошибка при маршализации JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при выполнении POST запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Ошибка при распаковке JSON:", err)
		return
	}

	fmt.Println("Ответ от модели:", response.Response)
}
