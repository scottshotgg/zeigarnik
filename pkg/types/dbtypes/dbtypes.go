package dbtypes

import (
	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

const (
	CREATED = "CREATED"
	QUEUED  = "QUEUED"
	FIRED   = "FIRED"
	MISSED  = "MISSED"
)

// Reminder is used for the persistence layer
type (
	Reminder struct {
		ID      string
		Created int64
		Message string
		To      string
		Status  string
		When    int64
		TypeOf  string
		WarnAt  []int64
	}
)

// ToDB maps an API reminder to a DB reminder
func ToDB(r *reminder.Reminder) *Reminder {
	return &Reminder{
		ID:      r.GetId(),
		Created: r.GetCreated(),
		When:    r.GetWhen(),
		Message: r.GetMessage(),
		Status:  r.GetStatus().String(),
		To:      r.GetTo(),
	}
}
