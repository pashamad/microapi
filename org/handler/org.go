package handler

import (
	"context"
	"github.com/micro/micro/v3/service/errors"
	"github.com/pashamad/microapi/org/models"
	"gorm.io/gorm"

	log "github.com/micro/micro/v3/service/logger"

	org "github.com/pashamad/microapi/org/proto"
)

var (
	ErrMissingUUID = errors.BadRequest("MISSING_UUID", "Missing UUID")
)

type Org struct {
	DB *gorm.DB
}

func (e *Org) Lookup(ctx context.Context, req *org.LookupRequest, rsp *org.LookupResponse) error {
	log.Info("Received Org.Lookup request")

	if len(req.Tin) == 0 {
		return ErrMissingUUID
	}

	db := e.DB
	var entity models.Entity

	err := db.Model(&models.Entity{}).Where(&models.Entity{TIN: req.Tin}).First(&entity).Error
	if err != nil {
		log.Error("Failed to read org.Entity from db")
		return err
	}
	log.Debugf("Entity UUID: %s", entity.UUID)

	rsp.Entity = entity.Serialize()

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Org) Stream(ctx context.Context, req *org.StreamingRequest, stream org.Org_StreamStream) error {
	log.Infof("Received Org.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&org.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Org) PingPong(ctx context.Context, stream org.Org_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&org.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
