package todo

import "time"

// Todo model
type Todo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar" json:"title"`
	Detail    string    `gorm:"type:varchar" json:"detail"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
