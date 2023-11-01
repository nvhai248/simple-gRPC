package main

import (
	"context"
	"fmt"
	"log"
	simple_gprc "simple-gprc/proto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("...Hello Client...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()
	client := simple_gprc.NewClassRegistrationServiceClient(cc)
	request := &simple_gprc.ClassRegistrationStatRequest{ClassIds: []int32{1, 2, 3, 4}}

	resp, _ := client.GetClassRegistrationStat(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp.Result)
}
