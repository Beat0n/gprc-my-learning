package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ZzCoding530/gprc-my-learning/hellogrpc/hellogrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultContent = "2024届秋招好难啊,学长快来帮帮我!giraffealex@163.com"
	defaultName    = "无名小将"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	content = flag.String("content", defaultContent, "content to be translate")
	name    = flag.String("name", defaultName, "your name")
	gender  = flag.String("gender", "male", "male or female")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslateClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.TranslateToEnglish(ctx, &pb.Input{
		InputContent: *content,
		Author: &pb.Author{
			Name:   *name,
			Gender: *gender,
		},
	})
	if err != nil {
		log.Fatalf("could not translate to English: %v", err)
	}
	log.Printf("英文翻译结果: %s|||%s", r.GetCall(), r.GetOutputContent())

	r, err = c.TranslateToGermany(ctx, &pb.Input{
		InputContent: *content,
		Author: &pb.Author{
			Name:   *name,
			Gender: *gender,
		},
	})
	if err != nil {
		log.Fatalf("could not translate to English: %v", err)
	}
	log.Printf("德文翻译结果: %s|||%s", r.GetCall(), r.GetOutputContent())
}
