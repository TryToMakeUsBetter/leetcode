package servcer

import (
	"context"

	//导入我们在protos文件中定义的服务
	pb "leetcode/grpc/proto/helloworld"
)

// 定义一个结构体，作用是实现helloworld中的GreeterServer
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
