package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestData struct {
	Volumes []int `json:"volumes"`
	K       int   `json:"k"`
}

type ResponseData struct {
	Operations int `json:"operations"`
}

func main() {
	http.HandleFunc("/align", alignHandler)
	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}

func alignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var requestData RequestData
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Ошибка парсинга данных", http.StatusBadRequest)
		return
	}

	n := len(requestData.Volumes)

	if requestData.K > n {
		http.Error(w, "K превышает количество сосудов", http.StatusBadRequest)
		return
	}

	maxVolume := requestData.Volumes[0]
	for i := 1; i < requestData.K; i++ {
		if requestData.Volumes[i] > maxVolume {
			maxVolume = requestData.Volumes[i]
		}
	}

	operations := 0
	for i := 0; i < requestData.K; i++ {
		if requestData.Volumes[i] > maxVolume {
			http.Error(w, "Некоторые сосуды имеют объем больше максимального среди первых K", http.StatusBadRequest)
			return
		}
		operations += maxVolume - requestData.Volumes[i]
	}

	responseData := ResponseData{Operations: operations}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
