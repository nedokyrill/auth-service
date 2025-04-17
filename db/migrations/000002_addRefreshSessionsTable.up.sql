create table if not exists "refreshSessions"
(
    "id"           serial primary key,
    "userId"       uuid references users(id) on delete cascade,
    "refreshToken" uuid                     not null,
    "ip"           character varying(15)    not null,
    "expiresIn"    bigint                   not null,
    "createdAt"    timestamp with time zone not null default now()

)