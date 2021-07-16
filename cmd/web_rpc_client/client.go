package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	v1 "web_rpc/api/order/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewOrderClient(conn)

	if reply, err := client.CreateOrder(context.Background(), &v1.CreateOrderRequest{UserName: "We"}); err != nil {
		log.Fatal(err)
	} else {
		log.Println(reply)
	}

	if reply, err := client.GetOrder(context.Background(), &v1.GetOrderRequest{Id: 1}); err != nil {
		log.Fatal(err)
	} else {
		log.Println(reply)
	}

	if reply, err := client.ListOrder(context.Background(), &v1.ListOrderRequest{
		Query:    "",
		Sequence: "AMOUNT",
		By:       "DESC",
	}); err != nil {
		log.Fatal(err)
	} else {
		for i, v := range reply.Orders {
			log.Println(i, v)
		}
	}

	if reply, err := client.UpdateOrder(context.Background(), &v1.UpdateOrderRequest{
		Id:      2,
		Amount:  9.991111,
		Status:  "offline",
		FileUrl: "",
	}); err != nil {
		log.Fatal(err)
	} else {
		log.Println(reply)
	}
}

