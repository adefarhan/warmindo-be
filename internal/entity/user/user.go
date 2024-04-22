package user

import "time"

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phoneNumber"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
