package dbtypes

import (
	reminder "github.com/scottshotgg/zeigarnik/pkg/reminder/v1alpha1"
)

// Reminder is used for the persistence layer
type (
	Reminder struct {
		ID      string
		Created int64
		Moment  int64
		Message string
		Status  string
		To      string
	}
)

// ToDB maps an API reminder to a DB reminder
func ToDB(r *reminder.Reminder) *Reminder {
	return &Reminder{
		ID:      r.GetId(),
		Created: r.GetCreated(),
		Moment:  r.GetMoment(),
		Message: r.GetMessage(),
		Status:  r.GetStatus(),
		To:      r.GetTo(),
	}
}
