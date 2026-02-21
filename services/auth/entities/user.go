package enti

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"type:varchar(20);default:'customer'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Hook ของ GORM: ก่อนสร้าง User ให้ Gen UUID ให้เอง
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}