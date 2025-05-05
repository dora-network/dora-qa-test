package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"dora-dev-test/consumer"
	"dora-dev-test/data"
	"dora-dev-test/generator"
	"dora-dev-test/publisher"
	"dora-dev-test/redis"
	"dora-dev-test/service"

	api "dora-dev-test/api/v1"

	redisv9 "github.com/redis/go-redis/v9"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = 8090
)

func main() {
	tickCh := make(chan data.Tick)
	go generator.GenerateTick(context.Background(), tickCh)
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("incoming_prices"),
	)
	if err != nil {
		panic(err)
	}
	rdb := redisv9.NewClient(&redisv9.Options{
		Addr: "localhost:6379",
	})
	ds := redis.NewDataStore(rdb)
	con := consumer.NewConsumer(client, ds)
	con.Start(context.Background())
	pub := publisher.NewTickPublisher(client, kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil))
	pub.Start(context.Background(), tickCh, "incoming_prices")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	api.RegisterDoraDevTestServiceServer(grpcServer, service.NewService(ds))
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(lis))
}
