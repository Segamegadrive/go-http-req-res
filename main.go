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

func mongoConn() {
	clientoptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientoptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully, connected to mongo database")

}

func getMetrics() {
	url := "http://aio-2521:8041/v1/metric/1823fb2e-b3b8-4aa4-b1bc-d154a8704cd6/measures"
	token := "gAAAAABe1pvvNJfWYJXBf61Nhx6rM2Ag6L5uMmzkGmmrYKBbuKYA4YheRD7dWlnRr59zgzO6PuTN6CNfiT5FyyaaIkVAVX0koyBluoLkFziHZkGAq9qM4Hu4Iiv5ouM6THdj5x5bvvhBuR6pMgS4-J4GH3eoLn8PmXJgzD9XgMwJdBjO0SkkvDM"
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

func main() {
	mongoConn()
	getMetrics()
}
