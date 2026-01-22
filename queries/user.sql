-- name: GetUserByID :one
select * from app.user u
where u.id = $1;

-- name: GetAllUsers :many
select * from app.user u;

-- name: GetUserByEmail :one
select * from app.user u
where u.email = $1;

-- name: CreateUser :one
insert into app.user(name, email, password, birthday)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateUserById :one
update app.user set
	name = $2,
	email = $3,
	password  = $4,
	birthday = $5
where id = $1
returning *;

-- name: DeleteUserById :exec
delete from app.user u
where u.id = $1;
