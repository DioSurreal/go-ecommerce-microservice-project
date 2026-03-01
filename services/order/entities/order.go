package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Order struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID      `gorm:"type:uuid;index"`
	TotalAmount float64        `gorm:"type:decimal(12,2)"`
	Status      string         `gorm:"type:varchar(50)"`
	// เก็บเป็นก้อน JSONB ใน Postgres
	Items       datatypes.JSON `gorm:"type:jsonb"` 
	CreatedAt   time.Time
}