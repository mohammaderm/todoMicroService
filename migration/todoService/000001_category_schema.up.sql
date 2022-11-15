CREATE TABLE if NOT EXISTS category(
    id int AUTO_INCREMENT unique not null,
    title VARCHAR(64) unique NOT NULL,
    accountid INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);
