package rpcClient

import (
	"github.com/RaymondCode/simple-demo/douyinBasic"
	"google.golang.org/grpc"
)

var (
	BasicPbAddress = "localhost:50051"
	BasicPbConn    *grpc.ClientConn
	BasicClient    douyinBasic.BasicApiServiceClient
)

func InitRpcClient() {
	var err error
	BasicPbConn, err = grpc.Dial(BasicPbAddress, grpc.WithTransportCredentials)
	if err != nil {
		panic(err)
	}

	BasicClient = douyinBasic.NewBasicApiServiceClient(BasicPbConn)
}
