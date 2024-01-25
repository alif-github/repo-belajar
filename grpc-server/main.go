package main

import (
	"context"
	"encoding/json"
	pb "github.com/alif-github/grpc-server/student"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	log.Println("Incoming request for FindStudentByEmail...")
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	var (
		data []byte
		err  error
	)

	data, err = ioutil.ReadFile("data/datas.json")
	if err != nil {
		log.Fatalln("error in read file -> ", err.Error())
	}

	if err = json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error unmarshal data -> ", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	var (
		listen net.Listener
		err    error
	)

	listen, err = net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in listen -> ", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve grpc -> ", err.Error())
	}
}
