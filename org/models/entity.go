package models

import (
	"github.com/google/uuid"
	pb "github.com/pashamad/microapi/org/proto"
)

type Entity struct {
	UUID uuid.UUID `gorm:"column:uuid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `gorm:"column:name"`
	TIN  string    `gorm:"column:tin;uniqueIndex:org_entity_tin_uindex"`
}

type Tabler interface {
	TableName() string
}

func (Entity) TableName() string {
	return "organisations"
}

func (e *Entity) Serialize() *pb.Entity {
	return &pb.Entity{
		Uuid: e.UUID.String(),
		Name: e.Name,
		Tin:  e.TIN,
	}
}
