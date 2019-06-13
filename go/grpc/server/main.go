package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "sayhi/grpc"

	"google.golang.org/grpc"
)

type server struct {
	hi bool

	number int64
	name   string
	msg    string
}

func (s *server) SayHello(ctx context.Context, in *pb.Hello) (*pb.Null, error) {
	reply := &pb.Null{}
	s.number = in.GetNum()
	s.name = in.GetName()
	s.msg = in.GetMsg()
	fmt.Printf("%d: %s Say %s\n", s.number, s.name, s.msg)
	s.hi = true
	return reply, nil
}

func (s *server) SayBye(in *pb.Null, stream pb.Greeter_SayByeServer) error {
	for {
		if s.hi {
			bye := &pb.Bye{
				Msg: "Receive: " + s.msg,
			}
			if (s.number % 2) == 0 {
				bye.Num = pb.Bye_ONE
			} else {
				bye.Num = pb.Bye_TWO
			}
			if err := stream.Send(bye); err != nil {
				return err // Disconnect
			}
			s.hi = false
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10086")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, new(server))
	fmt.Println("Server start ...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
