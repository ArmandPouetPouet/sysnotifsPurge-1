package handlers

import (
	"account-api-user/helpers"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"sort"

	"strings"
)

const userID = "userId"
const ratioMin = 20

type indexStruct struct {
	Indices map[string]struct {
		Primaries string `json:"primaries"`
	} `json:"indices"`
}

type jsonr struct {
	Nodes struct {
		Fs struct {
			TotalInBytes     int `json:"total_in_bytes"`
			AvailableInBytes int `json:"available_in_bytes"`
		} `json:"fs"`
	} `json:"nodes"`
}

//CheckRatio function : get user information
func CheckRatio(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetUser")

	resp, err := http.Get("https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/_cluster/stats")
	if err != nil {
		// handle error
		helpers.SetResponse(w, http.StatusOK, "error")
		return

	}
	defer resp.Body.Close()

	body, err0 := ioutil.ReadAll(resp.Body)
	if err0 != nil {
		// handle error
		helpers.SetResponse(w, http.StatusOK, "error")
		return

	}
	var b = string(body)
	fmt.Println(b)
	var data jsonr

	json.Unmarshal(body, &data)

	var ratio = int(float64(data.Nodes.Fs.AvailableInBytes) / float64(data.Nodes.Fs.TotalInBytes) * 100)

	if ratio < ratioMin {
		purgeIndex(searchOldIndex())

	}

	s := fmt.Sprintf("TotalInBytes %d   AvailableInBytes %d ratio %d",
		data.Nodes.Fs.TotalInBytes, data.Nodes.Fs.AvailableInBytes, ratio)

	helpers.SetResponse(w, http.StatusOK, s)

}

func searchOldIndex() string {

	resp, err := http.Get("https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/_stats")
	if err != nil {
		// handle error

	}
	defer resp.Body.Close()

	body, err0 := ioutil.ReadAll(resp.Body)
	if err0 != nil {
		// handle error

	}

	var b = string(body)
	fmt.Println(b)
	var data indexStruct

	json.Unmarshal(body, &data)

	var keys []string
	for key := range data.Indices {
		if !strings.Contains(key, "kibana") {
			keys = append(keys, key)
		}

	}
	sort.Strings(keys)

	fmt.Println("old index " + keys[0])

	return keys[0]
	//	fmt.Printf("Results: %v\n", data.Indices)

}

func purgeIndex(indexName string) {

	fmt.Println("purgeIndex " + indexName)
	req, err := http.NewRequest("DELETE", "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/"+indexName, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error

	}
	defer resp.Body.Close()

}
