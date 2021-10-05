CREATE TABLE IF NOT EXISTS articles(
    id serial PRIMARY KEY,
    title VARCHAR (150) NOT NULL,
    body text NOT NULL,
    posted_user_id integer NOT NULL,
    release_flg boolean default false,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;