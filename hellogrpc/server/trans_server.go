package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ZzCoding530/gprc-my-learning/hellogrpc/hellogrpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTranslateServer
}

func (t *server) TranslateToEnglish(ctx context.Context, input *pb.Input) (*pb.OutPut, error) {
	log.Println("TranslateToEnglish正在被调用")

	thisOutPut := &pb.OutPut{}

	thisOutPut.OutputContent = input.InputContent

	if input.Author.Gender == "male" {
		thisOutPut.Call = fmt.Sprintf("%s %s", input.Author.Name, "先生")
	} else {
		thisOutPut.Call = fmt.Sprintf("%s %s", input.Author.Name, "女士")
	}
	fmt.Println(thisOutPut.GetCall(), "我虽然是负责翻译英语，但是【", thisOutPut.GetOutputContent(), "】我也没听过啊，不知道怎么翻译成英文。")
	return &pb.OutPut{
		OutputContent: thisOutPut.OutputContent,
		Call:          thisOutPut.Call,
	}, nil
}

func (t *server) TranslateToGermany(ctx context.Context, input *pb.Input) (*pb.OutPut, error) {
	log.Println("TranslateToGermany正在被调用")

	thisOutPut := &pb.OutPut{}

	thisOutPut.OutputContent = input.InputContent

	if input.Author.Gender == "male" {
		thisOutPut.Call = fmt.Sprintf("%s %s", input.Author.Name, "先生")
	} else {
		thisOutPut.Call = fmt.Sprintf("%s %s", input.Author.Name, "女士")
	}
	fmt.Println(thisOutPut.GetCall(), "我虽然是负责翻译德语，但是【", thisOutPut.GetOutputContent(), "】这句话太难了，我刚学习德语一天啊。")
	return &pb.OutPut{
		OutputContent: thisOutPut.OutputContent,
		Call:          thisOutPut.Call,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslateServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
