CREATE TABLE IF NOT EXISTS advertisement
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    pictures    TEXT[] NOT NULL,
    price       INT NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


