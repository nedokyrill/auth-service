create extension if not exists "pgcrypto";

create table if not exists "users"
(
    "id" uuid primary key default gen_random_uuid(),
    "name" varchar(100) not null,
    "email" varchar(255) not null
);