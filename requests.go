package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
)

type Configuration struct {
    Port              int
    Base_url          string
    Token             string
}

func main() {
	// Define the URL
	var config = read_config()

    APIURL := config.Base_url + "search?q=Madonna&per_page=2" 
    // Define the Access token
    MyToken := "Bearer " + config.Token
    req, err := http.NewRequest(http.MethodGet, APIURL, nil)
    if err != nil {
        panic(err)
    }
    // add authorization header to the req
    req.Header.Add("Authorization", MyToken)
    resp := send_request(req)
    defer close_request(resp)
}

func read_config() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("Error when reading config:", err)
	}
	return configuration
}

func send_request(req *http.Request) *http.Response {
    // Send the request
    client := http.DefaultClient
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    return resp
}

func close_request(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    present_data(body)
}

func present_data(body []byte) {
	// handle the body data
    var data map[string]interface{}
    err2 := json.Unmarshal([]byte(body), &data)
    if err2 != nil {
        panic(err2)
    }
    // We are only interested in the response
    response_content := data["response"].(map[string]interface{})
	// We are only interested in the hits
    hits_content := response_content["hits"].(interface{})

    // Marshall the data into json for a nice representation
    json, err3 := json.MarshalIndent(hits_content, "", "    ")
    if err3 != nil {
    	panic(err3)
    }
    fmt.Println(string(json))


}
