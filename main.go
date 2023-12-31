package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("server run")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if message, ok := data["message"]; ok && reflect.TypeOf(message).Kind() == reflect.String {

		fmt.Println("Получено сообщение от клиента:", message)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Данные успешно приняты"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "400", "message": "Некорректное JSON-сообщение"})
	}
}
