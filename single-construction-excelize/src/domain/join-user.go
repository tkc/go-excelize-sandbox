package domain

import "time"

type JoinUser struct {
	ID        *int
	UserID    *int
	CreatedAt *time.Time
}
