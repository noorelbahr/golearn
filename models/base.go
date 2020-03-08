package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID        	uuid.UUID 	`gorm:"type:uuid;column:id;primary_key" json:"id"`
	CreatedAt 	time.Time	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt 	time.Time	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt 	*time.Time 	`sql:"index" gorm:"column:deleted_at" json:"-"`
}

/**
 * Custom Base Model Struct BeforeCreate function to set UUID
 */
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	// Generate UUID
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	return scope.SetColumn("id", uuid)
}