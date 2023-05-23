create table advertisement
(
    id          serial primary key,
    name        varchar(200) not null unique,
    description text         not null,
    pictures    serial,
    price       int          not null,
    created_at  timestamp    not null default now()
);