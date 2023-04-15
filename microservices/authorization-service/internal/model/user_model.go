package model

import(
	"time"
	"database/sql"
)
type User struct{
	ID        uint           `gorm:"primaryKey"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}