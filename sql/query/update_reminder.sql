-- name: UpdateReminder :one
update reminders set created = $1, msg = $2 where id = $2 returning *;