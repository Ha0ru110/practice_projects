package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://onboard.onboarddev.cloud.mobilaris.se:443/haohaotest"
	allData := []string{
		`{"longitude":1.1,"latitude":2}`,
		`{"longitude":3,"latitude":4}`,
		`{"longitude":5,"latitude":6}`,
		`{"longitude":7,"latitude":8}`,
		`{"longitude":9.0,"latitude":10}`,
		`{"longitude":-1,"latitude":-2}`,
		`{"longitude":-3,"latitude":-4}`,
		`{"longitude":-5,"latitude":-6}`,
		`{"longitude":-7,"latitude":-8}`,
		`{"longitude":-9.1,"latitude":-10}`,
	}
	for i := 0; i < 10; i++ {
		go sendJSONDataToServer(url, allData[i])
	}
	var whenFinished int
	for {
		whenFinished++
		if whenFinished == 10 {
			break
		}
	}

	fmt.Println("all info sent ")
}

func UploadData(gpsData string, sensorData string) {
	data := `{"longitude":1,"latitude":2,"gpsData",` + gpsData + `,"sensorData",` + sensorData + "}"
	url := "https://onboard.onboarddev.cloud.mobilaris.se:443/haohaotest"
	sendJSONDataToServer(url, data)
}
func sendJSONDataToServer(url string, content string) {
	jsonStr := []byte(content)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
}
