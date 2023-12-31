package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/protocol"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:50051"))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	protocol.RegisterAgentServer(server, &agentServer{})

	log.Println("Starting agent server...")
	log.Println(server.Serve(lis))
}

type agentServer struct {
	protocol.UnimplementedAgentServer
}

func (as *agentServer) Initialize(context.Context, *protocol.InitializeRequest) (*protocol.InitializeResponse, error) {
	return &protocol.InitializeResponse{
		Status: protocol.ResponseStatus_SUCCESS,
	}, nil
}

func (as *agentServer) EvaluateTx(ctx context.Context, txRequest *protocol.EvaluateTxRequest) (*protocol.EvaluateTxResponse, error) {
	return &protocol.EvaluateTxResponse{
		Status: protocol.ResponseStatus_SUCCESS,
	}, nil
}

func (as *agentServer) EvaluateBlock(ctx context.Context, req *protocol.EvaluateBlockRequest) (*protocol.EvaluateBlockResponse, error) {
	n, _ := hexutil.DecodeBig(req.Event.BlockNumber)
	resp := &protocol.EvaluateBlockResponse{
		Status: protocol.ResponseStatus_SUCCESS,
	}
	// publish a finding every 10 blocks
	if big.NewInt(0).Mod(n, big.NewInt(10)).Int64() == 0 {
		resp.Findings = append(resp.Findings, &protocol.Finding{
			Protocol:    "anonymous",
			Severity:    protocol.Finding_INFO,
			AlertId:     "HEARTBEAT",
			Name:        "Every 10 blocks",
			Description: n.String(),
		})
	}
	return resp, nil
}
