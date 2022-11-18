CREATE TABLE if NOT EXISTS todo(
    id serial UNIQUE PRIMARY key,
    title VARCHAR(64) NOT NULL,
    description TEXT,
    status BOOLEAN DEFAULT false,
    due_date TIMESTAMP NULL DEFAULT '1970-01-01 00:00:01',
    priority INT DEFAULT 0,
    categoryid INT NOT NULL,
    accountid INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
