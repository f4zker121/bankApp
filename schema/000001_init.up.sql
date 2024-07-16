CREATE TABLE users 
(
    id serial not null unique,
    balance FLOAT DEFAULT 0
);