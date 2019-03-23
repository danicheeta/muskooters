-- +goose Up
insert into users (username, passwd, role) VALUES ('admin', '$2a$10$7A/nol990ArOCdAxssX44.qSfcDpFPuYeRJ4A2gYkPemyzDrJzb9a', 'admin');