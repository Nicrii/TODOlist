package client

import (
	"context"
	"google.golang.org/grpc"
	"log"

	todo "../server/proto"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := todo.NewTodoServiceClient(conn)

	response, err := c.CreateTask(context.Background(), &todo.CreateRequest{Description: "new Task"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
