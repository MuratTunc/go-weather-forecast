package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *Config) Weather(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Message from weather service...",
		Data:    "DATA...",
	}

	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Turkey?unitGroup=metric&key=DTY5HP4XELL8CHTCGW65EXZ4G&contentType=json"

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
	payload.Message = string(body)
	payload.Data = string(body)

	app.writeJSON(w, http.StatusAccepted, payload)

}
