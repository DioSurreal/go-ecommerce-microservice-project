package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	CategoryID  int       `gorm:"index"`
	ImageURL    string    `gorm:"type:text"`
	// Relationship: 1 Product มี 1 Inventory
	Inventory Inventory `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Inventory struct {
	ProductID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	StockQuantity int       `gorm:"default:0"`
	Version       int       `gorm:"default:0"` // สำหรับ Optimistic Locking
	UpdatedAt time.Time
}
