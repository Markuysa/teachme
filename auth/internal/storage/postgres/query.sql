-- name: CreateUser :one
insert into "user" (id, name, email, created_at)
values ($1, $2, $3, $4)
    returning id, name, email, created_at;