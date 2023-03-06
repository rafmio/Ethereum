CREATE TABLE accounts (
    user_id serial PRIMARY KEY,
    accnumber VARCHAR (40) UNIQUE NOT NULL,
    password VARCHAR (256) NOT NULL
);