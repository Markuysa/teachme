
create table "user"
(
    id         uuid primary key,
    name       text      not null,
    email      text      not null,
    created_at timestamp not null
);

