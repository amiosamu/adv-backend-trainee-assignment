create table advertisement
(
    id          serial primary key,
    name        text      not null unique,
    description text      not null,
    pictures    text      not null,
    price       int       not null,
    created_at  timestamp not null default now()
);