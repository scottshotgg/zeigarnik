-- name: GetReminderByID :one
select * from reminders where id = $1;