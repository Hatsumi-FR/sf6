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

create table combos
(
    name         varchar not null,
    total_damage int,
    total_frame  int,
    difficulty   int,
    cost         int,
    id           serial
        constraint combos_pk
            primary key
);

create table inputs
(
    name            varchar,
    picture_url     varchar,
    frame           int,
    start_up_frame  int,
    active_frame    int,
    recovery_frame  int,
    damage          int,
    damage_on_block int,
    id              serial
        constraint inputs_pk
            primary key
);

create unique index inputs_id_uindex
    on inputs (id);


create unique index combos_id_uindex
    on combos (id);

create unique index characters_id_uindex
    on characters (id);

create unique index characters_name_uindex
    on characters (name);

