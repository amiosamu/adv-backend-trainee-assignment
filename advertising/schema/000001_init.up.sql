CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE advertising
(
    uuid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL DEFAULT '',
    description VARCHAR NOT NULL DEFAULT '',
    links TEXT[] NOT NULL,
    price NUMERIC NOT NULL
)
