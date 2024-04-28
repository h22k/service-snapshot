package model

import (
	"gorm.io/gorm"
)

type Snapshot struct {
	gorm.Model
	Url   string `gorm:"not null;uniqueIndex;unique"`
	Image []byte `gorm:"not null"`
}
