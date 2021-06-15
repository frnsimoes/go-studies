CREATE DATABASE devbook;
\c devbook;

CREATE TABLE usuarios(
        id int primary key,
        nome varchar(50) not null,
        nick varchar(50) not null unique,
        email varchar(50) not null unique,
        senha varchar(20) not null unique,
        criadoEm TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);