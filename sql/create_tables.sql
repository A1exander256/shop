CREATE TABLE users
(
    id serial PRIMARY KEY not null UNIQUE,
    firstname varchar(50) not null,
    surname varchar(50),
    middlename varchar(50),
    fio varchar(152) GENERATED ALWAYS AS (firstname || ' ' || surname || ' ' || middlename) STORED,
    sex varchar(5) not null CHECK (sex in ('man', 'woman')),
    age int CHECK (age > 0 AND age < 100)
);

CREATE TABLE orders
(
    id serial PRIMARY KEY not null UNIQUE,
    user_id int REFERENCES users (id) ON DELETE CASCADE not null,
    currency VARCHAR(3) NOT NULL CHECK( currency in ('BYN', 'EUR', 'USD')),
    total_cost NUMERIC(10, 2) DEFAULT 0 
);

CREATE TABLE products
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    left_in_stock INT NOT NULL DEFAULT 0
);

CREATE TABLE prices
(
    id serial PRIMARY KEY not null UNIQUE,
    product_id int REFERENCES products (id) ON DELETE CASCADE not null,
    currency VARCHAR(3) NOT NULL,
    price NUMERIC(10, 2) DEFAULT 0    
);

CREATE TABLE products_order
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    order_id int REFERENCES orders (id) ON DELETE CASCADE not null,
    product_id int REFERENCES products (id) ON DELETE CASCADE not null,
    count int CHECK (count > 0) DEFAULT 1  
);