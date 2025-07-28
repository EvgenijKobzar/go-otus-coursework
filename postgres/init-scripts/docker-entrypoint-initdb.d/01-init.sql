create schema if not exists movies_online;

create table if not exists movies_online.serials
(
    id        serial primary key,
    sort      integer,
    active    char(1) default 'Y',
    file_id int DEFAULT NULL,
    title varchar(50) DEFAULT NULL,
    production_period varchar(50) DEFAULT NULL,
    rating numeric DEFAULT NULL,
    quality varchar(50) DEFAULT NULL,
    duration numeric DEFAULT NULL,
    description varchar(150) DEFAULT NULL,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );

create table if not exists movies_online.seasons
(
    id        serial primary key,
    serial_id int references movies_online.serials(id),
    sort      integer,
    active    char(1) default 'Y',
    title varchar(50) DEFAULT NULL,
    created_by int default null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );

create table if not exists movies_online.episodes
(
    id        serial primary key,
    quality varchar(50) DEFAULT NULL,
    rating numeric DEFAULT NULL,
    production_period varchar(50) DEFAULT NULL,
    serial_id int references movies_online.serials(id),
    season_id int references movies_online.seasons(id),
    sort      integer,
    active    char(1) default 'Y',
    file_id int DEFAULT NULL,
    title varchar(50) DEFAULT NULL,
    duration numeric DEFAULT NULL,
    description varchar(150) DEFAULT NULL,
    created_by int default null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );

create table if not exists movies_online.accounts
(
    ID        serial primary key,
    FIRST_NAME varchar(150) DEFAULT NULL,
    LAST_NAME varchar(150) DEFAULT NULL,
    Name varchar(150) DEFAULT NULL,
    LOGIN varchar(150) DEFAULT NULL,
    PASSWORD varchar(150) DEFAULT NULL,
    CREATED_AT timestamp with time zone not null default now(),
    UPDATED_AT timestamp with time zone not null default now()
    );

alter table movies_online.seasons add moderated    boolean default false;
alter table movies_online.episodes add moderated    boolean default false;