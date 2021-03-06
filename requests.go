// Simple http requests package to be tested on Genius API
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
    "strings"
)


func main() {
    // Define the URL
    var config = read_config()
    cl_args := string(os.Args[1])
    artist_url := build_url(config, cl_args)
    // Define the Access token
    MyToken := "Bearer " + config.Token
    req, err := http.NewRequest(http.MethodGet, artist_url, nil)
    if err != nil {
        panic(err)
    }
    // add authorization header to the req
    req.Header.Add("Authorization", MyToken)
    resp := send_request(req)
    defer close_request(resp)
}

func build_url(config Configuration, cl_args string) string {
    // Make sure the full url will not have any spaces
    artist_replacement := strings.Replace(cl_args, " ", "%20", -1)
    my_per_page := string(config.Per_page)
    fmt.Printf(my_per_page)
    // concatenate into a full URL
    artist_url := config.Base_url + "search?q=" + artist_replacement + "&per_page=2" + ""//my_per_page
    return artist_url
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
