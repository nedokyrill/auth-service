create table if not exists "refreshes"
(
    "id"           serial primary key,
    "userId"       uuid references users (id) on delete cascade,
    "refreshToken" varchar(255)             not null,
    "ip"           varchar(15)              not null,
    "expiresIn"    bigint                   not null,
    "createdAt"    timestamp with time zone not null default now()

);
