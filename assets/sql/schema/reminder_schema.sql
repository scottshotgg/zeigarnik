create table reminders (
    id uuid unique,
    created integer not null,
    msg varchar not null,
    to varchar not null,
    status varchar not null,
    when varchar not null,
    type varchar not null,
    warnAt varchar not null,
);