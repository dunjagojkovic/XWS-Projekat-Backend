DROP TABLE IF EXISTS 'USER'

CREATE TABLE 'USER' (
    'id' SERIAL primary key,
    'name' varchar(255) NOT NULL,
    'surname' varchar(255) NOT NULL,
    'email' varchar(255) NOT NULL,
    'password' varchar(255) NOT NULL,
    'address' varchar(255) NOT NULL,
    'city' varchar(255) NOT NULL,
    'country' varchar(255) NOT NULL,
    'phone number' varchar(255) NOT NULL
);