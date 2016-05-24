package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	//Replace with your APPID
	appID := "APPID"

	var location string
	fmt.Print("Location : ")
	_, err := fmt.Scanln(&location)

	if err != nil {
		fmt.Println("Cannot get location")

	} else {
		var condition = new(Conditions)
		var apiURL = "http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + appID + "&units=metric"
		err := getJson(apiURL, condition)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("---------------------")
			fmt.Print("Temperature : ")
			fmt.Println(condition.Main.Temperature)
			fmt.Println("---------------------")
			fmt.Print("Status : ")
			fmt.Println(condition.Weather)
			fmt.Println("---------------------")
			fmt.Println("Have a nice day!")
			fmt.Println("---------------------")
		}

	}

}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Conditions struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

type Weather struct {
	Description string `json:"description"`
}

type Main struct {
	Temperature float32 `json:"temp"`
}
