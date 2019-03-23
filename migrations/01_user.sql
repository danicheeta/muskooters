-- +goose Up
create TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username varchar(40),
  passwd varchar(60),
  role varchar(10),
  UNIQUE(username)
);

-- +goose Down
drop TABLE users;