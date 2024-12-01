CREATE TABLE packages (
  id varchar(36) PRIMARY KEY NOT NULL,
  name varchar(255) NOT NULL,
  price int NOT NULL,
  description text,
  features varchar(255) NOT NULL,
  status bool NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at timestamp
)
