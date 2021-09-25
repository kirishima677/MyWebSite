CREATE TABLE IF NOT EXISTS users(
    user_id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;