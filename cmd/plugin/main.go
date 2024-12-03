package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/containerd/containerd/api/runtime/task/v3" // 替换为实际路径
	"google.golang.org/grpc"
)

// 定义一个结构体实现 Greeter 服务接口
type greeterServer struct {
	pb.UnimplementedTaskServer
}

// 实现 SayHello 方法
func (s *greeterServer) Create(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	log.Printf("Received: %s", req.ID)
	fmt.Println("Hello World-Create")
	os.WriteFile("/tmp/hello-world", []byte("Hello World-Create"), 0644)
	return &pb.CreateTaskResponse{Pid: 1033}, nil
}

func main() {
	lis, err := net.Listen("unix", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &greeterServer{})

	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//package main
//
//import (
//	context "context"
//	"fmt"
//	taskAPI "github.com/containerd/containerd/api/runtime/task/v3"
//	"net"
//	"os"
//
//	"google.golang.org/grpc"
//)
//
//type HelloWorldTaskService struct {
//	taskAPI.UnimplementedTaskServer // 实现默认接口
//}
//
//func (h *HelloWorldTaskService) Create(ctx context.Context, req *taskAPI.CreateTaskRequest) (*taskAPI.CreateTaskResponse, error) {
//	os.WriteFile("/tmp/hello-world", []byte("Hello World-Create"), 0644)
//	fmt.Println("Hello World-Create")
//	return &taskAPI.CreateTaskResponse{
//		Pid: 123,
//	}, nil
//}
//
//func main() {
//	// Provide a unix address to listen to, this will be the `address`
//	// in the `proxy_plugin` configuration.
//	// The root will be used to store the snapshots.
//	if len(os.Args) < 3 {
//		fmt.Printf("invalid args: usage: %s <unix addr> <root>\n", os.Args[0])
//		os.Exit(1)
//	}
//
//	// Create a gRPC server
//	rpc := grpc.NewServer()
//	taskAPI.RegisterTaskServer(rpc, &HelloWorldTaskService{})
//
//	// 监听 Unix Socket
//	l, err := net.Listen("unix", os.Args[1])
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("Hello World Plugin is running...")
//	if err := rpc.Serve(l); err != nil {
//		panic(err)
//	}
//
//}
