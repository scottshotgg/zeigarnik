-- name: DeleteReminder :one
delete from reminders where id = $1 returning *;