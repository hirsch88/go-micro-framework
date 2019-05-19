package models

import "time"

type Identifiable struct {
	ID        uint64     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
}

type UseTimestamp struct {
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
