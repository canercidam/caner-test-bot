package main

import (
	"context"
	"testing"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/require"
)

func TestEvaluateBlock(t *testing.T) {
	r := require.New(t)

	agentServer := &agentServer{}

	blockResp, err := agentServer.EvaluateBlock(context.Background(), &protocol.EvaluateBlockRequest{
		Event: &protocol.BlockEvent{
			BlockNumber: "0x1",
		},
	})
	r.NoError(err)

	r.Nil(blockResp.Findings)

	blockResp, err = agentServer.EvaluateBlock(context.Background(), &protocol.EvaluateBlockRequest{
		Event: &protocol.BlockEvent{
			BlockNumber: "0xa",
		},
	})
	r.NoError(err)

	r.NotNil(blockResp.Findings)
	r.Len(blockResp.Findings, 1)
}
