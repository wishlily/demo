package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	pb "sayhi/grpc"

	"google.golang.org/grpc"
)

func main() {
	wait := make(chan int, 1)

	conn, err := grpc.Dial("localhost:10086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "Robyn"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	msg := "Hello"
	if len(os.Args) > 2 {
		msg = os.Args[2]
	}
	num := int64(120)
	if len(os.Args) > 3 {
		n, _ := strconv.Atoi(os.Args[3])
		num = int64(n)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SayHello(ctx, &pb.Hello{Name: name, Msg: msg, Num: num})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s Say %s\n", num, name, msg)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		stream, err := c.SayBye(ctx, &pb.Null{})
		if err != nil {
			panic(err)
		}
		for {
			bye, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Printf("Say Bye %v %v\n", bye.GetNum(), bye.GetMsg())
			break
		}
		wait <- 0
	}()

	<-wait
}
