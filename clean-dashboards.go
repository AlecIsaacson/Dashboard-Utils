package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

//Define the dashboard JSON
type nrDashboardJSONStruct struct {
	Dashboard struct {
		//ID         int       `json:"id"`
		Title string `json:"title"`
		Icon  string `json:"icon"`
		//CreatedAt  time.Time `json:"created_at"`
		//UpdatedAt  time.Time `json:"updated_at"`
		Visibility string `json:"visibility"`
		Editable   string `json:"editable"`
		//UIURL      string    `json:"ui_url"`
		//APIURL     string    `json:"api_url"`
		//OwnerEmail string    `json:"owner_email"`
		Metadata struct {
			Version int `json:"version"`
		} `json:"metadata"`
		Widgets []struct {
			Visualization string `json:"visualization"`
			Layout        struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Row    int `json:"row"`
				Column int `json:"column"`
			} `json:"layout"`
			//WidgetID  int `json:"widget_id"`
			//AccountID int `json:"account_id"`
			Data []struct {
				Nrql string `json:"nrql"`
			} `json:"data"`
			Presentation struct {
				Title string      `json:"title"`
				Notes interface{} `json:"notes"`
				//DrilldownDashboardID int         `json:"drilldown_dashboard_id"`
			} `json:"presentation"`
		} `json:"widgets"`
		Filter struct {
			EventTypes []string `json:"event_types"`
			Attributes []string `json:"attributes"`
		} `json:"filter"`
	} `json:"dashboard"`
}

//Write each dashboard JSON definition to a file.
func writeFile(nrDashboardDefnPretty []byte, dashboardID string) {
	ioutil.WriteFile(dashboardID+".json", nrDashboardDefnPretty, 0644)
}

func main() {
	fmt.Println("NR Dashboard Cleaner v1.0")
	nrDashFileName := flag.String("file", "", "Exported dashboard JSON to clean.")
	flag.Parse()

	nrDashFile, err := ioutil.ReadFile(*nrDashFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Unmarshal the dashboard list into a struct
	var nrDashboardJSON nrDashboardJSONStruct
	if err := json.Unmarshal(nrDashFile, &nrDashboardJSON); err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", nrDashboardJSON)

	//Put the clean dashboard into a new file.
	fmt.Println("Cleaning :", nrDashboardJSON.Dashboard.Title)
	nrCleanJSON, err := json.MarshalIndent(nrDashboardJSON, "", "\t")

	// var nrDashboardDefnPretty bytes.Buffer
	// json.Indent(&nrDashboardDefnPretty, nrDashboardJSON, "", "\t")
	//fmt.Println(string(nrDashboardDefnPretty.Bytes()))
	writeFile(nrCleanJSON, nrDashboardJSON.Dashboard.Title)
}
