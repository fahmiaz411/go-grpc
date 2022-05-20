package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/fahmiaz411/go-grpc/student"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)

	if err != nil {
		log.Fatalln("error get student by email", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption
	
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":9000", opts...)
	if err != nil {
		log.Fatalln("error dial", err.Error())
	}

	defer conn.Close()

	client := pb.NewDataStudentClient(conn)

	getDataStudentByEmail(client, "fahmi@gmail.com")
	getDataStudentByEmail(client, "ega@gmail.com")
	
}