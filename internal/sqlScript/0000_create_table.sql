-- +migrate Up

-- Category

create table characters
(
    id            serial
        constraint characters_pk
            primary key,
    name          varchar not null,
    fighting_type int     not null,
    health        int,
    difficulty    int,
    mobility      int,
    range         int,
    power         int
);

create unique index characters_id_uindex
    on characters (id);

create unique index characters_name_uindex
    on characters (name);

