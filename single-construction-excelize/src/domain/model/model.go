package model

import "time"

type Excel struct {
	UserID           int
	UserName         string
	StartedDate      *time.Time
	StartedDatetime  *time.Time
	EndedDatetime    *time.Time
	ConstructionName string
	Memo             string
	Address          string
	SalesUserName    string
}

type JoinUser struct {
	ID        *int
	UserID    *int
	CreatedAt *time.Time
}
