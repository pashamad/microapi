package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	auth "github.com/pashamad/microapi/auth/proto"
)

type Auth struct{}

func (e *Auth) LoginApple(ctx context.Context, req *auth.LoginAppleRequest, rsp *auth.LoginAppleResponse) error {
	log.Info("Received Auth.LoginApple request")
	log.Infof("Token: %s", req.GetToken())
	log.Infof("Context vars: %v", ctx)
	rsp.Token = "JWT_AUTH_TOKEN_STUB"
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Auth) Stream(_ context.Context, req *auth.StreamingRequest, stream auth.Auth_StreamStream) error {
	log.Infof("Received Auth.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&auth.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Auth) PingPong(_ context.Context, stream auth.Auth_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&auth.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
