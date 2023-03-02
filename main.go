package main
    
import (
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
)
            
func main() {    
    url := "https://sa-east-1.aws.data.mongodb-api.com/app/data-vrbpu/endpoint/data/v1"
    method := "POST"
       
    payload := strings.NewReader(`{
        "collection":"collection",
        "database":"golangAPI",
        "dataSource":"golangAPI",
        "projection": {"_id": 1}
    }`)
            
    client := &http.Client{}

    req, err := http.NewRequest(method, url, payload)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Access-Control-Request-Headers", "*")
    req.Header.Add("api-key", "7Cxe7M29c2GcZljAsz02QpVSXTtwQM3MMah6suK2OLD3G2VOTRyrHr8Z7dhQZ4h7")
            
    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()
            
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}
