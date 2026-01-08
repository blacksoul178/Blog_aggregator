-- +goose up
CREATE TABLE users(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name varchar(100) not null UNIQUE
);
-- +goose down
DROP TABLE users;