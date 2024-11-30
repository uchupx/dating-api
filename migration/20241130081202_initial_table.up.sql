CREATE TABLE users(
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  client_app_id VARCHAR(36) NOT NULL,
  name VARCHAR(255) NULL,
  gender VARCHAR(1) NULL,
  address TEXT NULL,
  password TEXT NOT NULL,
  email varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NULL
);

CREATE TABLE refresh_tokens (
    id varchar(36) NOT NULL PRIMARY KEY,
    user_id varchar(36) NOT NULL,
    client_app_id varchar(36) NOT NULL,
    token text NOT NULL,
    expired_at DATETIME NOT NULL
  );

CREATE TABLE client_apps(
    id varchar(36) NOT NULL UNIQUE,
    `key` text NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    secret text NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
