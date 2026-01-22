create extension if not exists "uuid-ossp";

create schema app;

create table app.user(
	id                serial primary key,
	name              varchar(200) not null,
	email             varchar(200) not null unique,
	password          varchar(200) not null,
	birthday          date         not null,
	created_at        timestamp    not null default now()
);
