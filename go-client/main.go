package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "os"
    "strconv"
)

var (
        
        RPS := 10
)

func make_request(wg *sync.WaitGroup) (interface{}, error) { // Using sync.Waitgroup to wait for goroutines to finish
        serverUrl := "http://webserver/"

        c := http.Client{Timeout: time.Duration(1) * time.Second}
        resp, err := c.Get(serverUrl)
        if err != nil {
            fmt.Printf("Error %s", err)
            return
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Printf("Body : %s", body)
        wg.Done()
        return l, nil
}


func main() {
    
    if os.Getenv("RPS") != ""
      RPS = strconv.Atoi(os.Getenv("RPS"))

    fmt.Printf("Concurrent requests set to: %d", RPS)
    
    var wg sync.WaitGroup // New wait group
    wg.Add(RPS) // Using two goroutines
    go make_request(&wg)
    wg.Wait()


}