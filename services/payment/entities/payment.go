package entities

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID       uuid.UUID `gorm:"foreignKey::OrderID"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	Method        string    `gorm:"type:varchar(255);not null"`
	TransactionID string    `gorm:"type:varchar(255);not null"`
	Status        string    `gorm:"type:varchar(10);not null"`
	PaidAt        time.Time 
}
