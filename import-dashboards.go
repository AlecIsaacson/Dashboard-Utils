package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("NR Dashboard Importer v1.0")
	nrAPI := flag.String("apikey", "", "New Relic account API key")
	nrDashDefn := flag.String("defn", "", "New Relic dashboard definition JSON file")
	flag.Parse()

	nrDashJSON, err := os.Open(*nrDashDefn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer nrDashJSON.Close()

	nrBaseURL := "https://api.newrelic.com/v2/dashboards.json"

	client := &http.Client{}
	req, err := http.NewRequest("POST", nrBaseURL, nrDashJSON)
	req.Header.Set("X-Api-Key", *nrAPI)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	fmt.Println("Posting dashboard:", *nrDashDefn)
	fmt.Println("New Relic Response:", resp.Status)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("New Relic error")
		fmt.Println(resp)
	}

}
