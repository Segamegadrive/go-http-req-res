package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoConn()
	getMetrics()
}

// MongoDB connection
func mongoConn() {
	// Create mongo client and connect
	clientoptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientoptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check a client can connect
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully, connected to mongo database")
}

// Retrieve data from API call
func getMetrics() {

	url := "http://aio-2521:8041/v1/metric/1823fb2e-b3b8-4aa4-b1bc-d154a8704cd6/measures"
	token := "gAAAAABe2QL-gDs7h8y50GRB7_dv0lvS0nDPgCHvwgL3VyMV4av85u0MqDw1tm6QbH132Hpu-db97VQFcX7ly80WDJwerPY3XtILK_Hw1DFCgxrxhh4x8nggpIAHsdxdIxa7rmCVLqbgTnz2w3MedrC0P2NZGpnIZYBlmZIi5IbCPYsAkEQI0IQ"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("HTTP error %s\n", err)
	} else {
		req.Header.Add("x-auth-token", token)
		// Client required for HTTP headers
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(data))
		}
	}
}
