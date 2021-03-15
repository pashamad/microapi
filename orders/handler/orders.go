package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	orders "github.com/pashamad/microapi/orders/proto"
)

type Orders struct{}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Orders) Stream(ctx context.Context, req *orders.StreamingRequest, stream orders.Orders_StreamStream) error {
	log.Infof("Received Orders.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&orders.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Orders) PingPong(ctx context.Context, stream orders.Orders_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&orders.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
