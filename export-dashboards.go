package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type nrDashboardListStruct struct {
	Dashboards []struct {
		ID         int         `json:"id"`
		Title      string      `json:"title"`
		Icon       string      `json:"icon"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Visibility string      `json:"visibility"`
		Editable   string      `json:"editable"`
		UIURL      string      `json:"ui_url"`
		APIURL     string      `json:"api_url"`
		OwnerEmail string      `json:"owner_email"`
		Filter     interface{} `json:"filter"`
	} `json:"dashboards"`
}

//Make a NR Dashboard API call, returning the result as a byte string.
func getURL(urlToGet string, nrAPI string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlToGet, nil)
	req.Header.Set("X-Api-Key", nrAPI)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	fmt.Println("New Relic Response:", resp.Status)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("New Relic error")
		fmt.Println(resp)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)

	return response
}

//Write each dashboard JSON definition to a file.
func writeFile(nrDashboardDefnPretty []byte, dashboardID string) {
	ioutil.WriteFile(dashboardID+".json", nrDashboardDefnPretty, 0644)
}

func main() {
	fmt.Println("NR Dashboard Importer v1.0")
	nrAPI := flag.String("apikey", "", "New Relic account API key")
	nrDashID := flag.String("dashID", "", "A New Relic dashboard ID")
	flag.Parse()

	//Define the APIs base URL
	nrBaseURL := "https://api.newrelic.com/v2/"

	if *nrDashID == "" {
		//Get the list of all defined dashboards.
		nrDashboardListJSON := getURL(nrBaseURL+"dashboards.json", *nrAPI)

		// fmt.Printf("%s\n", nrDashboardListJSON)

		//Unmarshal the dashboard list into a struct
		var nrDashboardList nrDashboardListStruct
		if err := json.Unmarshal(nrDashboardListJSON, &nrDashboardList); err != nil {
			panic(err)
		}

		// fmt.Printf("%+v\n", nrDashboardList)

		//For each dashboard ID in the struct, get the dashboard definition, make it pretty, and dump it to a file.
		for _, dashboard := range nrDashboardList.Dashboards {
			fmt.Println("Getting :", dashboard.ID)
			nrDashboardDefn := getURL(nrBaseURL+"dashboards/"+strconv.Itoa(dashboard.ID)+".json", *nrAPI)
			var nrDashboardDefnPretty bytes.Buffer
			json.Indent(&nrDashboardDefnPretty, nrDashboardDefn, "", "\t")
			//fmt.Println(string(nrDashboardDefnPretty.Bytes()))
			writeFile(nrDashboardDefnPretty.Bytes(), strconv.Itoa(dashboard.ID))
		}
	} else {
		// We're just getting a single dashboard definition.
		fmt.Println("Getting :", *nrDashID)
		nrDashboardDefn := getURL(nrBaseURL+"dashboards/"+*nrDashID+".json", *nrAPI)
		var nrDashboardDefnPretty bytes.Buffer
		json.Indent(&nrDashboardDefnPretty, nrDashboardDefn, "", "\t")
		//fmt.Println(string(nrDashboardDefnPretty.Bytes()))
		writeFile(nrDashboardDefnPretty.Bytes(), *nrDashID)
	}

}
