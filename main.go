package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://aio-2521:8041/v1/metric/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("HTTP error %s\n", err)
	} else {
		req.Header.Add("x-auth-token", "gAAAAABe1laWSJy5aAjQZ5HgFCSapt1aOZAwDBAspPWTrb7SnzuK0AOOHUYqxSZFQd2VWIMWCrWYPUYGM-zLbvuXO30pq77nFPDbI4TFLFemwDao0fNLdCNQsSa52nPL5ZhyMwopArp0A3GS1mP0fHxNvo2Wle94kK5eRTafB_JkJ2ZcXHt8Pww")
		res, _ := http.DefaultClient.Do(req)
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}
