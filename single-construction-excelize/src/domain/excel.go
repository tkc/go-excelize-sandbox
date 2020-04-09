package domain

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
