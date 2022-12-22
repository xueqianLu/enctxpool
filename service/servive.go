package service

import (
	"context"
	"fmt"
	encservicetype "github.com/xueqianLu/enctxpool/protocol/generate/encservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type EncService struct {
	encservicetype.UnimplementedEncServiceServer
}

func (s *EncService) HealthCheck(ctx context.Context, req *emptypb.Empty) (*encservicetype.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (s *EncService) AddTx(ctx context.Context, req *encservicetype.AddTxRequest) (*encservicetype.AddTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTx not implemented")
}
func (s *EncService) Status(ctx context.Context, req *encservicetype.StatusRequest) (*encservicetype.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (s *EncService) Reset(ctx context.Context, req *encservicetype.ResetRequest) (*encservicetype.ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (s *EncService) Pending(ctx context.Context, req *encservicetype.PendingRequest) (*encservicetype.PendingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pending not implemented")
}

func RegisterService(server *grpc.Server) {
	encservicetype.RegisterEncServiceServer(server, &EncService{})
}

func StartEncService() {
	lis, err := net.Listen("tcp", ":38000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	RegisterService(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

