package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"sync"

	pb "github.com/fahmiaz411/go-grpc/student"
	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("server/data/datas.json")
	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error unmarshal", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()

	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("error listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error serve grpc", err.Error())
	}
}