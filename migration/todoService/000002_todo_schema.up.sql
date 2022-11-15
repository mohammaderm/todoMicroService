CREATE TABLE if NOT EXISTS todo(
    id int AUTO_INCREMENT unique not null,
    title VARCHAR(64) NOT NULL,
    description TEXT,
    status TINYINT DEFAULT 0,
    due_date TIMESTAMP NULL DEFAULT '1970-01-01 00:00:01',
    priority INT DEFAULT 0,
    categoryid INT NOT NULL,
    accountid INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);
