CREATE TABLE if NOT EXISTS category(
    id serial UNIQUE PRIMARY key,
    title VARCHAR(64) UNIQUE NOT NULL,
    accountid INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
