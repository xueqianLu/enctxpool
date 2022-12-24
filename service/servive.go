package service

import (
	"context"
	"errors"
	"fmt"
	corecmn "github.com/xueqianLu/enctxpool/core/common"
	"github.com/xueqianLu/enctxpool/node"
	encservicetype "github.com/xueqianLu/enctxpool/protocol/generate/encservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type EncService struct {
	n *node.Node
	encservicetype.UnimplementedEncServiceServer
}

func (s *EncService) HealthCheck(ctx context.Context, req *emptypb.Empty) (*encservicetype.HealthCheckResponse, error) {
	res := new(encservicetype.HealthCheckResponse)
	res.Status = true
	return res, nil
}
func (s *EncService) AddTx(ctx context.Context, req *encservicetype.AddTxRequest) (*encservicetype.AddTxResponse, error) {
	res := new(encservicetype.AddTxResponse)
	tx := corecmn.ParseTxData(req.Txdata)
	if tx != nil {
		err := s.n.TxPool().AddLocal(tx)
		if err != nil {
			res.Error = err.Error()
		}
		res.Txhash = tx.Hash().Bytes()
	} else {
		res.Error = errors.New("invalid tx data").Error()
	}

	return res,nil
}

func (s *EncService) Status(ctx context.Context, req *encservicetype.StatusRequest) (*encservicetype.StatusResponse, error) {
	pending, queue := s.n.TxPool().Stats()
	res := new(encservicetype.StatusResponse)
	res.Pending = uint64(pending)
	res.Queue = uint64(queue)
	return res, nil
}
func (s *EncService) Reset(ctx context.Context, req *encservicetype.ResetRequest) (*encservicetype.ResetResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (s *EncService) Pending(ctx context.Context, req *encservicetype.PendingRequest) (*encservicetype.PendingResponse, error) {
	alltxs := s.n.TxPool().Pending(false)
	res := new(encservicetype.PendingResponse)
	res.Txs = make([][]byte, 0)
	for _, txs := range alltxs {
		for _, tx := range txs {
			d, _ := tx.MarshalBinary()
			res.Txs = append(res.Txs, d)
		}
	}
	return res, nil
}

func RegisterService(server *grpc.Server, n *node.Node) {
	s := new(EncService)
	s.n = n
	encservicetype.RegisterEncServiceServer(server, s)
}

func StartEncService(n *node.Node) {
	lis, err := net.Listen("tcp", ":38000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	RegisterService(s, n)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

