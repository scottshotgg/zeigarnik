create table reminders (
    id uuid unique,
    created integer not null,
    msg varchar not null
);