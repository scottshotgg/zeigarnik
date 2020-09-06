-- name: CreateReminder :one
insert into reminders (id, created, msg) values ($1, $2, $3) returning *;