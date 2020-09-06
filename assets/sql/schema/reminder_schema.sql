create table reminders (
    id uuid unique,
    created bigint not null,
    msg varchar not null,
    recip varchar not null,
    rstatus varchar not null,
    atWhen varchar not null,
    typeOf varchar not null,
    warnAt bigint[] not null
);