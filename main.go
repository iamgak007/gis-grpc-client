package main

import (
	"context"
	"log"
	"time"

	pb "github.com/iamgak007/gis-grpc-client/corredor" // Change this to your path

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//change on the basis of services
	client := pb.NewMadinaGisServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//method
	req := &pb.UTMRequest{
		X: "56195",
		Y: "270585",
	}
	res, err := client.ConvertUTMToLatLon(ctx, req)
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}

	log.Println("lat lon")

	log.Println(res)

	req1 := &pb.Empty{}
	res1, err := client.GetSatelliteViewData(ctx, req1)
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}

	log.Println("satellite view data")
	log.Println(res1)
	// res, err = client.ConvertUTMToLatLon(ctx, req)
	// if err != nil {
	// 	log.Fatalf("could not register: %v", err)
	// }

	// log.Println(res)
	// log.Println(res)
}
