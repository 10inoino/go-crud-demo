-- +migrate Up
CREATE TABLE album(id text primary key, title text not null, artist text not null, price int not null);

-- +migrate Down
DROP TABLE album;