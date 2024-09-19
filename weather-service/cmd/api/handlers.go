package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Message from weather service...",
	}

	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/istanbul?unitGroup=metric&key=DTY5HP4XELL8CHTCGW65EXZ4G&contentType=json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("x-rapidapi-key", "DTY5HP4XELL8CHTCGW65EXZ4G")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(body))

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}
