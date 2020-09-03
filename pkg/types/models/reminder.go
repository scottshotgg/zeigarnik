package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	CREATED = "CREATED"
	QUEUED  = "QUEUED"
	FIRED   = "FIRED"
	MISSED  = "MISSED"
)

type (
	Reminder struct {
		ID      string
		Created int64
		When    int64
		Message string
		Status  string
		To      string
	}
)

func New(created, when int64, message, to string) *Reminder {
	return &Reminder{
		ID:      uuid.New().String(),
		Created: created,
		When:    when,
		Message: message,
		To:      to,
		Status:  CREATED,
	}
}

func (r *Reminder) IsAfter(t time.Duration) bool {
	var ttl = time.Duration(r.When - time.Now().Unix())

	if ttl > 0 && ttl < t {
		return true
	}

	return false
}

func (r *Reminder) CalcTTL() time.Duration {
	return time.Duration(r.When-time.Now().Unix()) * time.Second
}
