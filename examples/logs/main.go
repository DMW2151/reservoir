package main

import (
	"context"
	"fmt"
	"net"
	"time"

	grpc "google.golang.org/grpc"
	peer "google.golang.org/grpc/peer"

	reservoir "github.com/dmw2151/reservoir"
	logproto "github.com/dmw2151/reservoir/examples/proto/logs"

	reflection "google.golang.org/grpc/reflection"
)

const (
	logsHeldInSample = 3
	logSamplerSeed   = 2151
)

type LogsServer struct {
	logproto.UnimplementedLogsServer
}

// Submit -
func (l LogsServer) Submit(ctx context.Context, req *logproto.LogRequest) (*logproto.LogResponse, error) {
	return &logproto.LogResponse{Ok: true}, nil
}

// LogSamplerNode
type LogSamplerNode struct {
	r reservoir.reservoirSample[string]
}

// reservoirSamplerInterceptor
func (lsn *LogSamplerNode) reservoirSamplerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		h, err := handler(ctx, req)

		if msg, ok := req.(*logproto.LogRequest); ok {
			p, _ := peer.FromContext(ctx)
			stored := lsn.r.ReadSample(msg.Msg)
			fmt.Printf("Stored (%t): [%s]\t%s\t%s\n", stored, time.Now(), p.Addr.String(), msg.Msg)
		}

		return h, err
	}
}

func main() {

	logSamplerNode := LogSamplerNode{
		r: reservoir.NewreservoirSample[string](
			logsHeldInSample, logSamplerSeed,
		),
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(logSamplerNode.reservoirSamplerInterceptor()),
	)
	reflection.Register(srv)

	logproto.RegisterLogsServer(srv, LogsServer{})

	// test...
	// grpcurl -d '{"msg": "hello world"}' -plaintext localhost:2151 logs.Logs/Submit
	lis, _ := net.Listen("tcp", "0.0.0.0:2151")
	srv.Serve(lis)
}
