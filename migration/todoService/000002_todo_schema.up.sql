CREATE TABLE if NOT EXISTS todo(
    id int AUTO_INCREMENT unique not null,
    title VARCHAR(64) NOT NULL,
    description TEXT,
    Status TINYINT,
    due_date TIMESTAMP,
    priority INT,
    categoryid INT NOT NULL,
    accountid INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);
