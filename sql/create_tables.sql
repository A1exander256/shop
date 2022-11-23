CREATE TABLE users
(
    uuid serial PRIMARY KEY,
    firstname varchar(50) not null,
    surname varchar(50) not null,
    middlename varchar(50) not null,
    fio varchar(152) not null GENERATED ALWAYS AS (firstname || ' ' || surname || ' ' || middlename) STORED,
    sex varchar(5) not null CHECK (sex in ('man', 'woman')),
    age int not null CHECK (age > 0 AND age < 100)
);