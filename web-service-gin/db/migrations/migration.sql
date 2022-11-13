-- +migrate Up
CREATE TABLE album(id int, title text, artist text, price int);

-- +migrate Down
DROP TABLE album;