// Code generated by sqlc. DO NOT EDIT.
// source: get_reminder_by_id.sql

package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getReminderByID = `-- name: GetReminderByID :one
select id, created, msg, recip, rstatus, atwhen, typeof, warnat from reminders where id = $1
`

func (q *Queries) GetReminderByID(ctx context.Context, id uuid.UUID) (Reminder, error) {
	row := q.queryRow(ctx, q.getReminderByIDStmt, getReminderByID, id)
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
