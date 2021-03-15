package handler

import (
	"context"
	"receipt/parser"

	log "github.com/micro/micro/v3/service/logger"

	receipt "receipt/proto"
)

// todo: error definitions

type Receipt struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Receipt) Call(ctx context.Context, req *receipt.Message, rsp *receipt.Response) error {
	log.Info("Received Receipt.Call request")
	rsp.Result = "ok"
	return nil
}

func (e *Receipt) Scan(ctx context.Context, req *receipt.ScanRequest, rsp *receipt.ScanResponse) error {
	log.Info("Received Receipt.Scan request")
	log.Debug("DEBUG message")
	log.Trace("TRACE message")
	log.Info("INFO message")
	rsp.Amount = 0
	meta, err := parser.Parse(req.Data)
	if err != nil {
		return err
	}
	log.Info("Parsed code meta, ", meta)
	// restore order
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Receipt) Stream(ctx context.Context, req *receipt.StreamingRequest, stream receipt.Receipt_StreamStream) error {
	log.Infof("Received Receipt.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&receipt.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Receipt) PingPong(ctx context.Context, stream receipt.Receipt_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&receipt.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
