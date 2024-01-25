package main

import (
	"context"
	"fmt"
	pb "github.com/alif-github/grpc-client/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error when get student by email -> ", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial -> ", err.Error())
	}

	defer conn.Close()
	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "kuswandi@gmail.com")
}
