package model

import "time"

type User struct {
	ID        string
	Email     string
	Password  string
	Verified  bool
	Role      string
	CreatedAt time.Time
}
