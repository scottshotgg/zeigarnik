// Code generated by sqlc. DO NOT EDIT.
// source: update_reminder.sql

package sql

import (
	"context"

	"github.com/lib/pq"
)

const updateReminder = `-- name: UpdateReminder :one
update reminders set created = $1, msg = $2 where id = $2 returning id, created, msg, recip, rstatus, atwhen, typeof, warnat
`

type UpdateReminderParams struct {
	Created int64  `json:"created"`
	Msg     string `json:"msg"`
}

func (q *Queries) UpdateReminder(ctx context.Context, arg UpdateReminderParams) (Reminder, error) {
	row := q.queryRow(ctx, q.updateReminderStmt, updateReminder, arg.Created, arg.Msg)
	var i Reminder
	err := row.Scan(
		&i.ID,
		&i.Created,
		&i.Msg,
		&i.Recip,
		&i.Rstatus,
		&i.Atwhen,
		&i.Typeof,
		pq.Array(&i.Warnat),
	)
	return i, err
}
