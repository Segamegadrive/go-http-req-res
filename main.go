package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://aio-2521:8041/v1/metric/1823fb2e-b3b8-4aa4-b1bc-d154a8704cd6/measures"
	token := "gAAAAABe1nEOwXdKY1bQ7-cuoPkMjuqvTG8_8hI9trCMK-MLG7jQsssQPBjlwRwapevfdoWkVAdT7jsSuGbOY02mHRgS73QFfvYjool8YqRKPV04Jlj2Z6T8RJS84sPyWcnBwHU7qMBFAPaNzLbOsb2WBbFv992yEmzV06Fz1OuzykUof2N5qD0"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("HTTP error %s\n", err)
	} else {
		req.Header.Add("x-auth-token", token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(data))
		}
	}

}
