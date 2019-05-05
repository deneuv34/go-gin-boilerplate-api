package users

import "time"

type Users struct {
	ID        uint   `gorm:"primary_key;unique_index;AUTO_INCREMENT" json:"id"`
	FullName  string `json:"full_name"`
	Email     string `gorm:"unique_index" json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u Users) TableName() string {
	return "users"
}
