create table catalog.product
(
    id          bigint auto_increment
        primary key,
    name        varchar(255) not null,
    code        varchar(100) not null,
    description text null,
    qty         int     default 0 null,
    active      boolean default false,
    deleted     boolean default false,
    constraint code
        unique (code)
);