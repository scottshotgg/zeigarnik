// Code generated by sqlc. DO NOT EDIT.

package sql

import (
	"github.com/google/uuid"
)

type Reminder struct {
	ID      uuid.UUID `json:"id"`
	Created int64     `json:"created"`
	Msg     string    `json:"msg"`
	Recip   string    `json:"recip"`
	Rstatus string    `json:"rstatus"`
	Atwhen  string    `json:"atwhen"`
	Typeof  string    `json:"typeof"`
	Warnat  []int64   `json:"warnat"`
}
