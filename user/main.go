package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/guiluizmaia/tcc-grpc-golang/user/pb"
	"google.golang.org/grpc"
)

func main() {
	amount := 205000
	bytesRequest := int(0)
	timeStart := time.Now()
	success := 0
	error := 0

	connection, err := grpc.Dial("server:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	for i := 0; i < amount; i++ {

		client := pb.NewUserServiceClient(connection)

		req := &pb.User{
			Id:          uuid.New().String(),
			Name:        "nameFake",
			LastName:    "lastNameFake",
			Age:         "50",
			Document:    "33333333333",
			Address:     "Rua X",
			Nationality: "Nationality Fake",
			MotherName:  "MotherNameFake",
			FatherName:  "FatherName",
			Gender:      "Gender Fake",
			Birthday:    "12/07/2000",
			Email:       "test@test.com",
		}

		res, err := client.CreateUser(context.Background(), req)

		if err != nil {
			error += 1
			continue
		}

		if res != nil {
			success += 1
		}

		bytesRequest = req.GetLen()

	}
	timeFinish := time.Now()
	timeofRequests := timeFinish.Sub(timeStart)

	log.Printf("Amount of requests: %d", amount)
	log.Printf("Bytes of one request: %d", bytesRequest)
	log.Printf("Time of all requests: %s", timeofRequests)
	log.Printf("Errors: %d", error)
	log.Printf("Success: %d", success)
}
