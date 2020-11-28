package user

import "time"

// User model
type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar" json:"name"`
	Email        string    `gorm:"type:varchar" json:"email"`
	PasswordHash string    `gorm:"type:varchar"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
}
