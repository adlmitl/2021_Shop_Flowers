create table if not exists flower."user"
(
    id       uuid        not null
    primary key,
    login    varchar(45) not null,
    password varchar(45) not null
    );

alter table flower."user"
    owner to postgres;

create table if not exists flower.flower
(
    id    uuid           not null
    primary key,
    name  varchar(85)    not null,
    price numeric(10, 2) not null
    );

alter table flower.flower
    owner to postgres;

