package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	amount := 70000
	bytesRequest := int64(0)
	timeStart := time.Now()
	success := 0
	error := 0

	for i := 0; i < amount; i++ {
		bodySend, _ := json.Marshal(map[string]string{
			"id":          uuid.New().String(),
			"name":        "nameFake",
			"lastName":    "lastNameFake",
			"age":         "50",
			"document":    "33333333333",
			"address":     "Rua X",
			"nationality": "Nationality Fake",
			"motherName":  "MotherNameFake",
			"fatherName":  "FatherName",
			"gender":      "Gender Fake",
			"birthday":    "12/07/2000",
			"email":       "test@test.com",
		})
		payload := bytes.NewBuffer(bodySend)

		resp, err := http.Post("http://server:8080/", "application/json", payload)

		if err != nil {
			error += 1
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)

		bytesRequest = resp.ContentLength

		defer resp.Body.Close()
		if body != nil {
			success += 1
		}
	}
	timeFinish := time.Now()
	timeofRequests := timeFinish.Sub(timeStart)

	log.Printf("Amount of requests: %d", amount)
	log.Printf("Bytes of one request: %d", bytesRequest)
	log.Printf("Time of all requests: %s", timeofRequests)
	log.Printf("Errors: %d", error)
	log.Printf("Success: %d", success)
}
