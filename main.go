package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

type weatherStatus struct {
	Status      string `json:"main"`
	Description string `json:"description"`
}

type weatherCurrent struct {
	Temparature float64         `json:"temp"`
	Humidity    int64           `json:"humidity"`
	Weather     []weatherStatus `json:"weather"`
}

type weather struct {
	Lat      float64          `json:"lat"`
	Lon      float64          `json:"lon"`
	Timezone string           `json:"timezone"`
	Current  weatherCurrent   `json:"current"`
	Hourly   []weatherCurrent `json:"hourly"`
}

func main() {
	appID := os.Getenv("APP_ID")
	if appID == "" {
		fmt.Println("you need to provide env variable env APP_ID")
		return
	}

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	link := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=49.6823986&lon=18.3232279&exclude=minutely&appid=%s&units=metric", appID)
	resp, err := httpClient.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var scrapedData weather
	err = json.Unmarshal(body, &scrapedData)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := template.ParseFiles("./template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, scrapedData)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(buf.Bytes())
	})

	srv := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
