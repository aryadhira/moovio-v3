package utils

import (
	"encoding/json"
	"net/http"
)

func APICall(url string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var client = &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := WebResponse{}
	response.Status = status
	response.Message = message
	response.Data = data

	return json.NewEncoder(w).Encode(response)
}
