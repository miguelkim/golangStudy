package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET",
		"http://www.dplus.company", nil)
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", body)
}
