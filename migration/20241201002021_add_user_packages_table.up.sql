CREATE TABLE user_packages (
  id varchar(36) PRIMARY KEY NOT NULL,
  user_id varchar(36) NOT NULL,
  feature varchar(255) NOT NULL,
  status bool NOT NULL,
  valid_until DATE NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

